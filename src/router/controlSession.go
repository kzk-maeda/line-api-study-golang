package router

import (
	"fmt"
	"router/session"
)

func controlSession(user_id string, next_question string) (session.Session, error) {
	// SessionTableから対象userIdのSessionを取得
	user_session, err := session.GetSession(user_id)
	fmt.Println("Session : ", user_session, " Error : ", err)

	// まだSessionレコードがない場合、新たにSessionを生成
	if err != nil {
		user_session, err = session.SetSession(user_id)
		return user_session, err
	}

	// TODO:SessionレコードのDataが古い場合、新たにSessionを生成
	
	// Sessionレコードが存在する場合、Questionを更新
	user_session, err = session.UpdateSession(user_id, next_question)
	return user_session, err
}

func controlTextSession(user_id string) (session.Session, error) {
	// SessionTableから対象userIdのSessionを取得
	user_session, err := session.GetSession(user_id)
	fmt.Println("Session : ", user_session, " Error : ", err)

	// まだSessionレコードがない場合、新たにSessionを生成
	if err != nil {
		user_session, err = session.SetSession(user_id)
		return user_session, err
	}

	// TODO:SessionレコードのDataが古い場合、新たにSessionを生成

	// Sessionレコードが存在する場合、Questionを更新
	// 現在のQuestionIdから、次のIDを算出して値を格納
	current_question_id := user_session.QuestionId
	next_question_id := "1"
	switch current_question_id {
	case "3-1":
		next_question_id = "3-2"
	case "3-2":
		next_question_id = "3-3"
	case "3-3":
		next_question_id = "4"
	case "6-6":
		next_question_id = "6-7"
	default:
		next_question_id = "1"
	}
	user_session, err = session.UpdateSession(user_id, next_question_id)
	return user_session, err
}