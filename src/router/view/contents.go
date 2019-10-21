package view

import (
// "encoding/json"
// "fmt"
)

type ActionComponent struct {
	Label string `json:"label"`
	Data  string `json:"data"`
}

type ButtonAction struct {
	Type        string `json:"type"`
	Label       string `json:"label"`
	Data        string `json:"data"`
	DisplayText string `json:"displayText"`
}

// Public Methods
func CreateQuestion(next_question string) string {
	// define
	question_text := "Question Default"
	body_type := "button"
	actionComponents := []ActionComponent{
		ActionComponent{
			Label: "selection1",
			Data: createAnswerData(
				"deefault", "selection_default", "2-1",
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
				Label: "はい",
				Data: createAnswerData(
					"1", "yes", "2-1",
				),
			},
			ActionComponent{
				Label: "いいえ",
				Data: createAnswerData(
					"1", "no", "3-1",
				),
			},
		}
	// 2-1）どのくらいの期間妊活していますか？
	case "2-1":
		question_text = "どのくらいの期間妊活していますか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "1年以内",
				Data: createAnswerData(
					"2-1", "1", "2-2",
				),
			},
			ActionComponent{
				Label: "2年以内",
				Data: createAnswerData(
					"2-1", "2", "2-2",
				),
			},
			ActionComponent{
				Label: "3-4年",
				Data: createAnswerData(
					"2-1", "3-4", "2-2",
				),
			},
			ActionComponent{
				Label: "5-6年",
				Data: createAnswerData(
					"2-1", "5-6", "2-2",
				),
			},
			ActionComponent{
				Label: "7年以上",
				Data: createAnswerData(
					"2-1", "7", "2-2",
				),
			},
		}
	// 2-2）現在、不妊治療中ですか？
	case "2-2":
		question_text = "現在、不妊治療中ですか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "はい",
				Data: createAnswerData(
					"2-2", "yes", "2-3",
				),
			},
			ActionComponent{
				Label: "いいえ",
				Data: createAnswerData(
					"2-2", "no", "2-3",
				),
			},
		}
	// 2-3）配偶者は精液検査を受けましたか？
	case "2-3":
		question_text = "配偶者は精液検査を受けましたか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "はい",
				Data: createAnswerData(
					"2-3", "yes", "2-4",
				),
			},
			ActionComponent{
				Label: "いいえ",
				Data: createAnswerData(
					"2-3", "no", "3-1",
				),
			},
		}
	case "2-4":
		question_text = "精液検査での運動率結果を選択してください"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "60％以上",
				Data: createAnswerData(
					"2-4", "60", "3-1",
				),
			},
			ActionComponent{
				Label: "40-59％",
				Data: createAnswerData(
					"2-4", "40-59", "3-1",
				),
			},
			ActionComponent{
				Label: "20-39％",
				Data: createAnswerData(
					"2-4", "20-39", "3-1",
				),
			},
			ActionComponent{
				Label: "0-19％",
				Data: createAnswerData(
					"2-4", "0-19", "3-1",
				),
			},
			ActionComponent{
				Label: "不明",
				Data: createAnswerData(
					"2-4", "unknown", "3-1",
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
	case "3-3":
		question_text = "体重は何kgですか？"
		body_type = "text"
		textComponent = "半角数字で入力してください"
	case "4":
		question_text = "以前妊娠したことはありますか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "はい",
				Data: createAnswerData(
					"4", "yes", "5",
				),
			},
			ActionComponent{
				Label: "いいえ",
				Data: createAnswerData(
					"4", "no", "6-1",
				),
			},
		}
	case "5":
		question_text = "出産経験はありますか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "はい",
				Data: createAnswerData(
					"5", "yes", "6-1",
				),
			},
			ActionComponent{
				Label: "いいえ",
				Data: createAnswerData(
					"5", "no", "6-1",
				),
			},
		}
	case "6-1":
		question_text = "生理周期は順調ですか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "はい（26~35日周期）",
				Data: createAnswerData(
					"6-1", "yes", "6-2",
				),
			},
			ActionComponent{
				Label: "いいえ",
				Data: createAnswerData(
					"6-1", "no", "6-2",
				),
			},
		}
	case "6-2":
		question_text = "生理痛は重いですか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "はい",
				Data: createAnswerData(
					"6-2", "yes", "6-3",
				),
			},
			ActionComponent{
				Label: "いいえ",
				Data: createAnswerData(
					"6-2", "no", "6-3",
				),
			},
		}
	case "6-3":
		question_text = "喫煙しますか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "はい",
				Data: createAnswerData(
					"6-3", "yes", "6-4",
				),
			},
			ActionComponent{
				Label: "いいえ",
				Data: createAnswerData(
					"6-3", "no", "6-4",
				),
			},
		}
	case "6-4":
		question_text = "1日平均どのくらいお酒を飲みますか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "１杯以上",
				Data: createAnswerData(
					"6-3", "more_1", "6-5",
				),
			},
			ActionComponent{
				Label: "１杯未満",
				Data: createAnswerData(
					"6-3", "less_1", "6-5",
				),
			},
		}
	case "6-5":
		question_text = "AMH（卵巣年齢）血液検査を1年以内に受けましたか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "はい",
				Data: createAnswerData(
					"6-5", "yes", "6-6",
				),
			},
			ActionComponent{
				Label: "いいえ",
				Data: createAnswerData(
					"6-5", "no", "6-7",
				),
			},
		}
	case "6-6":
		question_text = "AMH（卵巣年齢）検査の数値を入力してください"
		body_type = "text"
		textComponent = "半角数字、小数点第二位まで入力してください"
	case "6-7":
		question_text = "以下に該当する症状はありますか？\n\n" +
			"➀子宮内膜症、多嚢胞性卵巣症候群、子宮因子(妊娠に影響する筋腫など)\n\n" +
			"➁排卵障害、卵管因子、子宮頸がん、乳がん\n\n" +
			"➂卵巣機能不全（POI）、40歳以前の早期閉経\n\n" +
			"➃その他妊娠に関わらないと診断された疾患、特になし\n"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "➀",
				Data: createAnswerData(
					"6-7", "level_1", "7-1",
				),
			},
			ActionComponent{
				Label: "➁",
				Data: createAnswerData(
					"6-7", "level_2", "7-1",
				),
			},
			ActionComponent{
				Label: "➂",
				Data: createAnswerData(
					"6-7", "level_3", "7-1",
				),
			},
			ActionComponent{
				Label: "➃",
				Data: createAnswerData(
					"6-7", "level_0", "7-1",
				),
			},
		}
	case "7-1":
		question_text = "友達や家族に妊活や妊娠のための健康について話したことはありますか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "はい",
				Data: createAnswerData(
					"7-1", "yes", "7-2",
				),
			},
			ActionComponent{
				Label: "いいえ",
				Data: createAnswerData(
					"7-1", "no", "7-2",
				),
			},
		}
	case "7-2":
		question_text = "利用規約に同意しますか？"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "はい",
				Data: createAnswerData(
					"7-2", "yes", "result",
				),
			},
			ActionComponent{
				Label: "いいえ",
				Data: createAnswerData(
					"7-2", "no", "result",
				),
			},
		}
	case "result":
		question_text = "診断結果"
		body_type = "text"
		textComponent = "あなたの妊娠力はランクBです"

	// default）
	case "default":
		question_text = "Questoin?"
		body_type = "button"
		actionComponents = []ActionComponent{
			ActionComponent{
				Label: "Label",
				Data: createAnswerData(
					"x", "answer", "next_question",
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
