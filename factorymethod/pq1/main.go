package main

import (
	"fmt"

	"example.com/factorymethod/notification"
)

func main() {
	notifier, err := notification.NewNotifier("push")
	if err != nil {
		fmt.Println(err)
	}
	message := notifier.Send("Hello World!")
	fmt.Println(message)
}
