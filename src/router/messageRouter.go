package router

import (
	"fmt"
	"router/view"
	"router/session"
)

func MessageRouter(message_text string) string {
	// session controll
	session := session.SetSession("user_id")
	fmt.Println(session.QuestionId)
	replyText := view.CreateQuestion("default")
	switch message_text {
	case "test":
		fmt.Println("this is test message")
		replyText = view.CreateQuestion("1")
	case "test2":
		fmt.Println("this is test message2")
		replyText = view.CreateQuestion("2")
	}
	return replyText
}
