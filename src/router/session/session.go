package session

import (
	"fmt"
	"time"

	// "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Session struct {
	UserId string `dynamo:"user_id"`
	QuestionId string `dynamo:"question_id"`
  CreatedTime time.Time `dynamo:"created_time"`
}

func SetSession(userId string) Session {
	// 最初にSessionをSetするfunction
	table := initialize()
	session := Session{UserId: userId, QuestionId: "1", CreatedTime: time.Now().UTC()}

	err := table.Put(session).Run()
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

	return session
}

func GetSession(userId string) Session {
	// userIdからSessionを取得するfunction
	table := initialize()
	session := Session{}

	err := table.Get("user_id", userId).All(&session)
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

	return session
}

func UpdateSession(userId string, nextQuestionId string) {
	// Sessionを更新するfunction
	table := initialize()

	err := table.Update("user_id", userId).Set("question_id", nextQuestionId).Run()
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

}

func initialize() dynamo.Table {
	ddb := dynamo.New(session.New())
	table := ddb.Table("session_table")

	return table
}