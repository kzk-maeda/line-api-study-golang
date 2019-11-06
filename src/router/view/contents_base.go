package view

import (
	"encoding/json"
)

type ButtonContent struct {
	Type   string       `json:"type"`
	Style  string       `json:"style"`
	Action ButtonAction `json:"action"`
}

type TextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func createTextContents(component string) interface{} {
	// textContents := TextContent{}
	// textContent.Type = "text"
	// textContent.Text = component

	body_content := map[string]interface{}{
		"type":       "box",
		"layout":     "vertical",
		"spacing":    "md",
		"paddingAll": "12px",
		"contents": []interface{}{
			map[string]interface{}{
				"type": "text",
				"text": component,
				"wrap": true,
			},
		},
	}

	return body_content
}

func createButtonContents(actionComponents []ActionComponent) interface{} {
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

	body_content := map[string]interface{}{
		"type":   "box",
		"layout": "vertical",
		"contents": []interface{}{
			map[string]interface{}{
				"type":     "box",
				"layout":   "vertical",
				"spacing":  "md",
				"contents": buttonContents,
				"flex":     1,
			},
		},
		"spacing":    "md",
		"paddingAll": "12px",
	}

	return body_content
}

func createAnswerData(question string, answer string, next_question string) string {
	data_string := "question=" + question +
		"&answer=" + answer +
		"&next_question=" + next_question
	return data_string
}

// JSONのBaseコンテンツを生成する関数
func createBaseContents(question_text string, contents interface{}) string {
	map_flex_contents := map[string]interface{}{
		"type": "bubble",
		"header": map[string]interface{}{
			"type":            "box",
			"layout":          "vertical",
			"backgroundColor": "#66cdaa",
			"contents": []interface{}{
				map[string]interface{}{
					"type": "text",
					"text": question_text,
					// "color":   "#ffffff",
					"align":   "start",
					"size":    "md",
					"gravity": "center",
					"wrap":    true,
				},
				map[string]interface{}{
					"type":   "box",
					"layout": "vertical",
					"contents": []interface{}{
						map[string]interface{}{
							"type":   "box",
							"layout": "vertical",
							"contents": []interface{}{
								map[string]interface{}{
									"type": "filler",
								},
							},
							"width":           "10%",
							"backgroundColor": "#f0f0f0",
							"height":          "6px",
						},
					},
					"backgroundColor": "#9FD8E36E",
					"height":          "6px",
					"margin":          "sm",
				},
			},
			// "backgroundColor": "#66cdaa",
			"paddingTop":    "19px",
			"paddingAll":    "12px",
			"paddingBottom": "16px",
		},
		"body": contents,
		"styles": map[string]interface{}{
			"footer": map[string]interface{}{
				"separator": false,
			},
		},
	}
	json_flex_contents, _ := json.Marshal(map_flex_contents)
	flex_contents := string(json_flex_contents)

	return flex_contents
}
