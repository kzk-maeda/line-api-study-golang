package view

import (
	"encoding/json"
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

type ButtonContent struct {
	Type string `json:"type"`
	Style string `json:"style"`
	Action ButtonAction `json:"action"`
}

// Public Methods
func CreateQuestion(next_question string) string {
	// define
	question_text := "Question Default"
	actionComponents := []ActionComponent{
		ActionComponent{
			Label:"selection1",
			Data:createAnswerData(
				"question_no_deefault", "selection_default", "2-1",
			),
		}, 
	}

	// each contents
	switch next_question {
	case "1":
		question_text = "Question No 1"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label:"selection1",
				Data:createAnswerData(
					"question_no_1", "selection1", "2-1",
				),
			}, 
			ActionComponent{
				Label:"selection2",
				Data:createAnswerData(
					"question_no_1", "selection2", "3-1",
				),
			},
		}
	case "2":
		question_text = "Question No 2"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label:"selection1",
				Data:createAnswerData(
					"question_no_2", "selection1", "2-1",
				),
			}, 
			ActionComponent{
				Label:"selection2",
				Data:createAnswerData(
					"question_no_2", "selection2", "3-1",
				),
			},
		}
	}
	
	buttonContents := createButtonContents(actionComponents)
	contents := createBaseContents(question_text, buttonContents)

	// fmt.Println("Send FlexMessage as Below.")
	// fmt.Println(contents)
	return contents
}


// Private Methods

func createButtonContents(actionComponents []ActionComponent) []ButtonContent {
	buttonContents := []ButtonContent{}
	for _, actionComponent := range actionComponents {
		buttonContent := ButtonContent{}
		buttonAction := ButtonAction{}

		buttonContent.Type = "button"
		buttonContent.Style = "secondary"

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
					// "color":   "#ffffff",
					"align":   "start",
					"size":    "md",
					"gravity": "center",
					"wrap": true,
				},
				map[string]interface{}{
					"type": "box",
					"layout": "vertical",
					"contents": []interface{}{
						map[string]interface{}{
							"type": "box",
							"layout": "vertical",
							"contents": []interface{}{
								map[string]interface{}{
									"type": "filler",
								},
							},
							"width": "10%",
							"backgroundColor": "#f0f0f0",
							"height": "6px",
						},
					},
					"backgroundColor": "#9FD8E36E",
					"height": "6px",
					"margin": "sm",
				},
			},
			"backgroundColor": "#DE5658",
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
	// fmt.Println(flex_contents)

	return flex_contents
}
