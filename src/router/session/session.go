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
	UserId string `dynamo:"user_id"`
	QuestionId string `dynamo:"question_id"`
  CreatedTime time.Time `dynamo:"created_time"`
}

func SetSession(userId string) (Session, error) {
	// 最初にSessionをSetするfunction
	// table := initialize()
	session := Session{UserId: userId, QuestionId: "1", CreatedTime: time.Now().UTC()}

	err := SessionTable.Put(session).Run()
	if err != nil {
		fmt.Println(err)
		// panic(err.Error())
	}

	return session, err
}

func GetSession(userId string) (Session, error) {
	// userIdからSessionを取得するfunction
	// table := initialize()
	session := Session{}

	err := SessionTable.Get("user_id", userId).One(&session)
	if err != nil {
		fmt.Println(err)
		// panic(err.Error())
	}

	return session, err
}

func UpdateSession(userId string, nextQuestionId string) (Session, error) {
	// Sessionを更新するfunction
	// table := initialize()

	err := SessionTable.Update("user_id", userId).
		Set("question_id", nextQuestionId).
		Set("created_time", time.Now().UTC()).
		Run()
	if err != nil {
		fmt.Println(err)
		// panic(err.Error())
	}
	session, err := GetSession(userId)
	return session, err
}

func init() {
	ddb := dynamo.New(session.New())
	SessionTable = ddb.Table("session_table")

	return
}