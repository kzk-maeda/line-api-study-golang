package main

import (
	"encoding/json"
	"fmt"
)

type ButtonContentsItems struct {
	LabelText    string
	DataText     string
	NextQuestion string
}

type ButtonActionComponent struct {
	Type        string
	Label       string
	Data        string
	DisplayText string
}

type ButtonComponent struct {
	Type   string
	Style  string
	Action map[string]ButtonActionComponent
}

func createButtonComponents(question_label string, button_contents_items ButtonContentsItems) []ButtonComponent {
	buttonComponent1 := ButtonComponent{}
	buttonComponent2 := ButtonComponent{}
	buttonComponent := []ButtonComponent{buttonComponent1, buttonComponent2}

	return buttonComponent
}

// JSONのBaseコンテンツを生成する関数
func CreateBaseContents(h_text string) string {
	map_flex_contents := map[string]interface{}{
		"type": "bubble",
		"header": map[string]interface{}{
			"type":   "box",
			"layout": "vertical",
			"contents": []interface{}{
				map[string]interface{}{
					"type":    "text",
					"text":    h_text,
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
					"contents": []interface{}{
						map[string]interface{}{
							"type":  "button",
							"style": "primary",
							"action": map[string]interface{}{
								"type":        "postback",
								"label":       "answer 1",
								"data":        "question=q1&answer=a1&next_question=q2",
								"displayText": "answer 1",
							},
						},
						map[string]interface{}{
							"type":  "button",
							"style": "primary",
							"action": map[string]interface{}{
								"type":        "postback",
								"label":       "answer 2",
								"data":        "question=q1&answer=a2&next_question=q2",
								"displayText": "answer 2",
							},
						},
					},
					// "contents": createButtonComponents("test question")
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
