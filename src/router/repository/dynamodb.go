package repository

import (
	"fmt"
	"time"

	// "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type User struct {
	UserId string `dynamo:"user_id"`
  CreatedTime time.Time `dynamo:"created_time"`
}

func TestCall()  {
	fmt.Println("test")
	table := initialize()

	u := User{UserId: "3", CreatedTime: time.Now().UTC()}
	fmt.Println(u)

	if err := table.Put(u).Run(); err != nil {
    fmt.Println("err")
    panic(err.Error())
  }

}

func initialize() dynamo.Table {
	ddb := dynamo.New(session.New())
	table := ddb.Table("data_table")

	return table
}