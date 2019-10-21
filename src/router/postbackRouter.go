package router

import (
	"fmt"
	"router/repository"
	"router/view"
	"strings"
)

type PostbackData struct {
	Question     string
	Answer       string
	NextQuestion string
}

func PostbackRouter(user_id string, postbackData string) string {
	data := initializePostbackData(postbackData)
	controlSession(user_id, data.NextQuestion)

	// 回答をTableに登録
	repository.RegisterData(user_id, data.Question, "", data.Answer)

	// SessionTableに登録されるNextQuestionとpostbackDataから得られるNextQuestionが同値なので、
	// SessionTableからNextQuestionの値を取得する処理はスキップ
	replyText := view.CreateQuestion(data.NextQuestion)

	// NextQuestionが"result"だった場合は結果計算処理に移動
	if data.NextQuestion == "result" {
		rank := calculateRank(user_id)
		BMI := calculateBMI(user_id)
		fmt.Println(BMI)
		replyText = view.CreateResult(data.NextQuestion, rank)
	}

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
			return_data.NextQuestion = value
		}
	}

	repository.TestCall()

	return return_data
}
