package router

import (
	"fmt"
	"router/view"
)

func MessageRouter(user_id string, message_text string) string {
	replyText := view.CreateQuestion("default")
	switch message_text {
	case "test":
		fmt.Println("this is test message")
		controlSession(user_id, "1")
		replyText = view.CreateQuestion("1")
	case "test2":
		fmt.Println("this is test message2")
		controlSession(user_id, "2")
		replyText = view.CreateQuestion("2")
	}
	return replyText
}
