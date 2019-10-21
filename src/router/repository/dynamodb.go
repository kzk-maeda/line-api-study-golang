package repository

import (
	"fmt"
	"time"

	// "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var DataTable dynamo.Table

type Answer struct {
	// QuestionNo string `dynamo:"question_no,omitempty"`
	Description string    `dynamo:"description,omitempty"`
	Answer      string    `dynamo:"answer,omitempty"`
	UpdatedTime time.Time `dynamo:"updated_time,omitempty"`
}

type User struct {
	UserId      string             `dynamo:"user_id"`
	Answers     map[string]*Answer `dynamo:"answers,omitempty"`
	CreatedTime time.Time          `dynamo:"created_time"`
}

func RegisterData(user_id string, question_no string, description string, answer string) (User, error) {
	// user_idのレコードを取得：ない場合作成
	data, err := getData(user_id)
	fmt.Println(data.Answers)
	if err != nil {
		data, err = setData(user_id)
	}

	// Dataレコードを追加
	data, err = updateData(user_id, question_no, description, answer)

	return data, err
}

func setData(user_id string) (User, error) {
	data := User{UserId: user_id, CreatedTime: time.Now().UTC()}
	err := DataTable.Put(data).Run()
	if err != nil {
		fmt.Println(err)
	}

	return data, err
}

func updateData(user_id string, question_no string, description string, answer string) (User, error) {
	// DBに格納する構造体の定義
	answer_rec := &Answer{
		// QuestionNo: question_no,
		Description: description,
		Answer:      answer,
		UpdatedTime: time.Now().UTC(),
	}
	// Answersカラムが存在しない場合、作成して定義した構造体を追加
	data, err := getData(user_id)
	answer_map := map[string]*Answer{}
	if data.Answers != nil {
		answer_map = data.Answers
	}
	answer_map[question_no] = answer_rec

	// DDBにはQuestionNoをキーとした構造体を格納
	qerr := DataTable.Update("user_id", user_id).
		Set("answers", answer_map).
		Run()
	if qerr != nil {
		fmt.Println(qerr)
	}
	data, err = getData(user_id)
	return data, err
}

func getData(user_id string) (User, error) {
	data := User{}

	err := DataTable.Get("user_id", user_id).One(&data)
	if err != nil {
		fmt.Println(err)
	}

	return data, err
}

func TestCall() {
	fmt.Println("test")

	u := User{UserId: "3", CreatedTime: time.Now().UTC()}
	fmt.Println(u)

	if err := DataTable.Put(u).Run(); err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

}

func init() {
	ddb := dynamo.New(session.New())
	DataTable = ddb.Table("data_table")

	return
}
