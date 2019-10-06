package main

import (
	"fmt"
)

func MessageRouter(message_text string) string {
	replyText := "default"
	switch message_text {
	case "test":
		fmt.Println("this is test message")
		replyText = "test message from router"
	case "test2":
		fmt.Println("this is test message2")
		replyText = "test message2 from router"
	}
	return replyText
}
