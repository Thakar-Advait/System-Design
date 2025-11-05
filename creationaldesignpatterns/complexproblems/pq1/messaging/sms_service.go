package messaging

import (
	"fmt"
	"sync"
	"time"
)

type smsService struct {
	createdAt time.Time
}

func (s *smsService) PlatformName() string { return "SMS" }

func (s *smsService) Send(m *Message) error {
	if m == nil || m.Platform() != "SMS" {
		return fmt.Errorf("sms service can only send sms messages")
	}
	// simulate send
	fmt.Println("[SMSService] Sending:", m.Summary())
	return nil
}

func (s *smsService) Schedule(m *Message, when time.Time) error {
	if m == nil || m.Platform() != "SMS" {
		return fmt.Errorf("sms service can only schedule sms messages")
	}
	delay := time.Until(when)
	if delay <= 0 {
		return s.Send(m)
	}
	go func(msg *Message, d time.Duration) {
		timer := time.NewTimer(d)
		<-timer.C
		_ = s.Send(msg)
	}(m, delay)
	fmt.Printf("[SMSService] Scheduled message for %s (in %v)\n", when.Format(time.RFC3339), delay)
	return nil
}

var smsOnce sync.Once
var smsSvc *smsService

func getSMSServiceInstance() PlatformService {
	smsOnce.Do(func() {
		smsSvc = &smsService{createdAt: time.Now()}
		fmt.Println("[SMSService] instance created")
	})
	return smsSvc
}
