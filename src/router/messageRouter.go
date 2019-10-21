package router

import (
	// "fmt"
	"router/view"
	// "router/session"
)

func MessageRouter(user_id string, message_text string) string {
	replyText := view.CreateQuestion("default")
	// 次の質問IDをSessionTableに格納
	user_session, _ := controlTextSession(user_id)
	// 次の質問IDを取得
	next_question := user_session.QuestionId
	// replyTextを生成
	replyText = view.CreateQuestion(next_question)
	
	return replyText
}
