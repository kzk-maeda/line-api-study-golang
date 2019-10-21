package view

import (
	// "encoding/json"
	// "fmt"
)

type ActionComponent struct {
	Label string `json:"label"`
	Data string `json:"data"`
}

type ButtonAction struct {
	Type string `json:"type"`
	Label string `json:"label"`
	Data string `json:"data"`
	DisplayText string `json:"displayText"`
}

// Public Methods
func CreateQuestion(next_question string) string {
	// define
	question_text := "Question Default"
	body_type := "button"
	actionComponents := []ActionComponent{
		ActionComponent{
			Label:"selection1",
			Data:createAnswerData(
				"question_no_deefault", "selection_default", "2-1",
			),
		}, 
	}
	textComponent := "textComponent"

	// each contents
	switch next_question {
	// 1) 現在結婚していますか？
	case "1":
		question_text = "現在結婚していますか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label:"はい",
				Data:createAnswerData(
					"question_no_1", "yes", "2-1",
				),
			}, 
			ActionComponent{
				Label:"いいえ",
				Data:createAnswerData(
					"question_no_1", "no", "3-1",
				),
			},
		}
	// 2-1）どのくらいの期間妊活していますか？
	case "2-1":
		question_text = "どのくらいの期間妊活していますか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label:"1年以内",
				Data:createAnswerData(
					"question_no_2-1", "1", "2-2",
				),
			}, 
			ActionComponent{
				Label:"2年以内",
				Data:createAnswerData(
					"question_no_2-1", "2", "2-2",
				),
			},
			ActionComponent{
				Label:"3-4年",
				Data:createAnswerData(
					"question_no_2-1", "3-4", "2-2",
				),
			},
			ActionComponent{
				Label:"5-6年",
				Data:createAnswerData(
					"question_no_2-1", "5-6", "2-2",
				),
			},
			ActionComponent{
				Label:"7年以上",
				Data:createAnswerData(
					"question_no_2-1", "7", "2-2",
				),
			},
		}
	// 2-2）現在、不妊治療中ですか？
	case "2-2":
		question_text = "現在、不妊治療中ですか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label:"はい",
				Data:createAnswerData(
					"question_no_2-2", "yes", "2-3",
				),
			},
			ActionComponent{
				Label:"いいえ",
				Data:createAnswerData(
					"question_no_2-2", "yes", "2-3",
				),
			},
		}
	// 2-3）配偶者は精液検査を受けましたか？
	case "2-3":
		question_text = "配偶者は精液検査を受けましたか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label:"はい",
				Data:createAnswerData(
					"question_no_2-3", "yes", "2-4",
				),
			},
			ActionComponent{
				Label:"いいえ",
				Data:createAnswerData(
					"question_no_2-3", "yes", "3-1",
				),
			},
		}
	case "2-4":
		question_text = "精液検査での運動率結果を選択してください"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label:"60％以上",
				Data:createAnswerData(
					"question_no_2-4", "60", "3-1",
				),
			},
			ActionComponent{
				Label:"40-59％",
				Data:createAnswerData(
					"question_no_2-4", "40-59", "3-1",
				),
			},
			ActionComponent{
				Label:"20-39％",
				Data:createAnswerData(
					"question_no_2-4", "20-39", "3-1",
				),
			},
			ActionComponent{
				Label:"0-19％",
				Data:createAnswerData(
					"question_no_2-4", "0-19", "3-1",
				),
			},
			ActionComponent{
				Label:"不明",
				Data:createAnswerData(
					"question_no_2-4", "unknown", "3-1",
				),
			},
		}
	case "3-1":
		question_text = "現在何歳ですか？"
		body_type = "text"
		textComponent = "半角数字で入力してください"
	case "3-2":
		question_text = "身長は何cmですか？"
		body_type = "text"
		textComponent = "半角数字で入力してください"
	
	// default）
	case "default":
		question_text = "Questoin?"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label:"Label",
				Data:createAnswerData(
					"question_no_x", "answer", "next_question",
				),
			},
		}
		textComponent = "textComponent"
	// from here
	}
	
	contents := ""
	switch body_type {
	case "button":
		bodyContents := createButtonContents(actionComponents)
		contents = createBaseContents(question_text, bodyContents)
	case "text":
		bodyContents := createTextContents(textComponent)
		contents = createBaseContents(question_text, bodyContents)
	default:
		bodyContents := createTextContents(textComponent)
		contents = createBaseContents(question_text, bodyContents)
	}

	return contents
}
