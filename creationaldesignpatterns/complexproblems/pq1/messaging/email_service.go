package messaging

import (
	"fmt"
	"sync"
	"time"
)

// emailService is the actual delivery manager for emails.
// This is a singleton: only one per process.
type emailService struct {
	// in real world: connection pool, client credentials, metrics, etc.
	createdAt time.Time
}

func (s *emailService) PlatformName() string { return "EMAIL" }

func (s *emailService) Send(m *Message) error {
	// validate message platform
	if m == nil || m.Platform() != "EMAIL" {
		return fmt.Errorf("email service can only send email messages")
	}
	// simulate send - in prod you'd call SMTP/API
	fmt.Println("[EmailService] Sending:", m.Summary())
	return nil
}

func (s *emailService) Schedule(m *Message, when time.Time) error {
	if m == nil || m.Platform() != "EMAIL" {
		return fmt.Errorf("email service can only schedule email messages")
	}
	// simplified scheduling: spawn goroutine to simulate scheduled send
	delay := time.Until(when)
	if delay <= 0 {
		// immediate
		return s.Send(m)
	}

	go func(msg *Message, d time.Duration) {
		// in production: use persistent job queue, retry, backoff, etc.
		timer := time.NewTimer(d)
		<-timer.C
		_ = s.Send(msg)
	}(m, delay)

	fmt.Printf("[EmailService] Scheduled message for %s (in %v)\n", when.Format(time.RFC3339), delay)
	return nil
}

/***** singleton wiring *****/
var emailOnce sync.Once
var emailSvc *emailService

func getEmailServiceInstance() PlatformService {
	emailOnce.Do(func() {
		emailSvc = &emailService{createdAt: time.Now()}
		fmt.Println("[EmailService] instance created")
	})
	return emailSvc
}
