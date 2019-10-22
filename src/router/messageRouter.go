package router

import (
	// "fmt"
	"router/repository"
	"router/session"
	"router/view"
)

func MessageRouter(user_id string, message_text string) string {
	replyText := view.CreateQuestion("default")
	// 診断ツールのスタート："診断スタート"を受け取ると無条件で初期状態にする
	if message_text == "診断スタート" {
		session.DeleteSession(user_id)
		repository.DeleteAnswerData(user_id)
	}
	// TODO: Validation
	validate_check()

	// 今の質問番号をSessionから取得
	user_session, _ := session.GetSession(user_id)
	question_no := user_session.QuestionId
	// 回答をTableに登録
	repository.RegisterData(user_id, question_no, "", message_text)

	// 次の質問IDをSessionTableに格納
	user_session, _ = controlTextSession(user_id)
	// 次の質問IDを取得
	next_question := user_session.QuestionId
	// replyTextを生成
	replyText = view.CreateQuestion(next_question)

	return replyText
}

// TODO:
func validate_check()  {
	
}