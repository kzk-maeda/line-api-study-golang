package session

import (
	"fmt"
	"time"

	// "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var SessionTable dynamo.Table

type Session struct {
	UserId      string    `dynamo:"user_id"`
	QuestionId  string    `dynamo:"question_id"`
	CreatedTime time.Time `dynamo:"created_time"`
}

func SetSession(user_id string) (Session, error) {
	// 最初にSessionをSetするfunction
	// table := initialize()
	session := Session{UserId: user_id, QuestionId: "1", CreatedTime: time.Now().UTC()}

	err := SessionTable.Put(session).Run()
	if err != nil {
		fmt.Println(err)
		// panic(err.Error())
	}

	return session, err
}

func GetSession(user_id string) (Session, error) {
	// user_idからSessionを取得するfunction
	// table := initialize()
	session := Session{}

	err := SessionTable.Get("user_id", user_id).One(&session)
	if err != nil {
		fmt.Println(err)
		// panic(err.Error())
	}

	return session, err
}

func UpdateSession(user_id string, next_question_id string) (Session, error) {
	// Sessionを更新するfunction
	// table := initialize()

	err := SessionTable.Update("user_id", user_id).
		Set("question_id", next_question_id).
		Set("created_time", time.Now().UTC()).
		Run()
	if err != nil {
		fmt.Println(err)
		// panic(err.Error())
	}
	session, err := GetSession(user_id)
	return session, err
}

func init() {
	ddb := dynamo.New(session.New())
	SessionTable = ddb.Table("session_table")

	return
}

func DeleteSession(user_id string) {
	err := SessionTable.Delete("user_id", user_id).Run()
	if err != nil {
		fmt.Println(err)
	}
}