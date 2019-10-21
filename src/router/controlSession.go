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