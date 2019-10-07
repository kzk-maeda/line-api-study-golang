package main

import (
	"encoding/json"
	"fmt"
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

type ButtonContent struct {
	Type string `json:"type"`
	Style string `json:"style"`
	Action ButtonAction `json:"action"`
}

// Public Methods
func CreateQuestion() string {
	question_text := "Question No 1"
	
	selection1 := ActionComponent{
		Label:"selection1",
		Data:createAnswerData(
			"question_no_1", "selection1", "2-1",
		),
	}
	selection2 := ActionComponent{
		Label:"selection2",
		Data:createAnswerData(
			"question_no_1", "selection2", "3-1",
		),
	}
	actionComponents := []ActionComponent{selection1, selection2}

	buttonContents := createButtonContents(actionComponents)

	contents := createBaseContents(question_text, buttonContents)

	return contents
}


// Private Methods

func createButtonContents(actionComponents []ActionComponent) []ButtonContent {
	buttonContents := []ButtonContent{}
	for _, actionComponent := range actionComponents {
		buttonContent := ButtonContent{}
		buttonAction := ButtonAction{}

		buttonContent.Type = "button"
		buttonContent.Style = "primary"

		buttonAction.Type = "postback"
		buttonAction.Label = actionComponent.Label
		buttonAction.Data = actionComponent.Data
		buttonAction.DisplayText = actionComponent.Label

		buttonContent.Action = buttonAction

		buttonContents = append(buttonContents, buttonContent)
	}

	return buttonContents
}

func createAnswerData(question string, answer string, next_question string) string {
	data_string := "question=" + question +
								 "&answer=" + answer +
								 "&next_question=" + next_question
	return data_string
}

// JSONのBaseコンテンツを生成する関数
func createBaseContents(question_text string, contents []ButtonContent) string {
	map_flex_contents := map[string]interface{}{
		"type": "bubble",
		"header": map[string]interface{}{
			"type":   "box",
			"layout": "vertical",
			"contents": []interface{}{
				map[string]interface{}{
					"type":    "text",
					"text":    question_text,
					"color":   "#ffffff",
					"align":   "start",
					"size":    "md",
					"gravity": "center",
				},
			},
			"backgroundColor": "#27ACB2",
			"paddingTop":      "19px",
			"paddingAll":      "12px",
			"paddingBottom":   "16px",
		},
		"body": map[string]interface{}{
			"type":   "box",
			"layout": "vertical",
			"contents": []interface{}{
				map[string]interface{}{
					"type":    "box",
					"layout":  "vertical",
					"spacing": "md",
					"contents": contents,
					"flex": 1,
				},
			},
			"spacing":    "md",
			"paddingAll": "12px",
		},
		"styles": map[string]interface{}{
			"footer": map[string]interface{}{
				"separator": false,
			},
		},
	}
	json_flex_contents, _ := json.Marshal(map_flex_contents)
	flex_contents := string(json_flex_contents)
	fmt.Println(flex_contents)

	return flex_contents
}
