package router

import (
	"fmt"
	"strings"
	"router/repository"
	"router/view"
)

type PostbackData struct {
	Question    string
	Answer      string
	NexQuestion string
}

func PostbackRouter(user_id string, postbackData string) string {
	data := initializePostbackData(postbackData)
	controlSession(user_id, data.NexQuestion)
	replyText := view.CreateQuestion(data.NexQuestion)

	return replyText
}

// postback.dataをsplitしてObject化
func initializePostbackData(postbackData string) PostbackData {
	fmt.Println(postbackData)
	splited_data := strings.Split(postbackData, "&")
	return_data := PostbackData{}
	for _, item := range splited_data {
		key := strings.Split(item, "=")[0]
		value := strings.Split(item, "=")[1]

		fmt.Println(key, value)

		switch key {
		case "question":
			return_data.Question = value
		case "answer":
			return_data.Answer = value
		case "next_question":
			return_data.NexQuestion = value
		}
	}

	repository.TestCall()

	return return_data
}
