package router

import (
	// "fmt"
	"router/repository"
	"router/session"
	"router/view"
	"strconv"
)

func MessageRouter(user_id string, message_text string) string {
	replyText := view.CreateQuestion("1")
	// 診断ツールのスタート："診断スタート"を受け取ると無条件で初期状態にする
	if message_text == "診断スタート" {
		session.DeleteSession(user_id)
		repository.DeleteAnswerData(user_id)
		// 次の質問IDをSessionTableに格納
		_, err := controlTextSession(user_id)
		if err != nil {
			return replyText
		}
		replyText := view.CreateQuestion("1")
		return replyText
	}
	// Get Session ID
	user_session, _ := session.GetSession(user_id)
	question_no := user_session.QuestionId
	next_question := question_no

	// validation
	is_valid := validate_check(user_id, message_text)
	if is_valid == true { // ValidationがOKのとき
		// 回答をTableに登録
		repository.RegisterData(user_id, question_no, "", message_text)
		// 次の質問IDをSessionTableに格納
		user_session, _ = controlTextSession(user_id)
		// 次の質問IDを取得
		next_question = user_session.QuestionId

	} else { // ValidationがFalseのとき
		// TODO: 同じ質問をもう一度送付
	}

	// replyTextを生成
	replyText = view.CreateQuestion(next_question)
	return replyText
}

// MessageEventで送付されたTextが正しいか確認
func validate_check(user_id string, message_text string) (is_valid bool) {
	// Sessionから現在の質問番号を取得
	sess, err := session.GetSession(user_id)
	if err != nil {
		is_valid = false
		return is_valid
	}
	question_id := sess.QuestionId
	// question_idで分岐
	switch question_id {
	case "3-1": // 年齢
		age, err := strconv.Atoi(message_text)
		if err != nil {
			return false
		}
		if 12 <= age && age <= 80 {
			return true
		} else {
			return false
		}

	case "3-2": // 身長
		height, err := strconv.Atoi(message_text)
		if err != nil {
			return false
		}
		if 90 <= height && height <= 250 {
			return true
		} else {
			return false
		}

	case "3-3": // 体重
		weight, err := strconv.Atoi(message_text)
		if err != nil {
			return false
		}
		if 20 <= weight && weight <= 200 {
			return true
		} else {
			return false
		}

	case "6-6": // AMH
		amh, err := strconv.ParseFloat(message_text, 64)
		if err != nil {
			return false
		}
		if 0.0 <= amh && amh <= 10.0 {
			return true
		} else {
			return false
		}

	}
	return
}
