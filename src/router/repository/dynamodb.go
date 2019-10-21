package repository

import (
	"fmt"
	"time"

	// "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var DataTable dynamo.Table


type User struct {
	UserId string `dynamo:"user_id"`
  CreatedTime time.Time `dynamo:"created_time"`
}

func TestCall()  {
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