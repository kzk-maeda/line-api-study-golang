package view

import (
	"fmt"
	"strconv"
	"encoding/json"
	"router/dto"
)

func CreateResult(answer_params dto.AnswerParams) string {
	// message定義を読み込み
	result_message := resultMessage()
	
	bubble_list := []interface{}{}

	// val init
	var header_text string
	var body_text string
	var body_text_list []string

	// rankの値から診断結果Bubbleを生成
	header_text = "診断結果：ランク" + answer_params.Rank
	body_text = ""
	body_text_list = []string{
		"あなたの妊娠力はランク" + answer_params.Rank + "です",
		"詳細は右にスワイプしてください",
	}
	// main_bubble_body_list := createBubbleBodyContents(main_body_text_list)
	// main_bubble_contents_item := createBubbleContentsItem("診断結果：ランク" + answer_params.Rank, main_bubble_body_list)
	bubble_list = append(bubble_list, joinContents(header_text, body_text_list))

	// rankの値から医師コメントBubbleを生成
	header_text = "医師からのコメント"
	switch answer_params.Rank {
		case "A": body_text = result_message.DoctorComment.A
		case "B": body_text = result_message.DoctorComment.B
		case "C": body_text = result_message.DoctorComment.C
		case "D": body_text = result_message.DoctorComment.D
		case "E": body_text = result_message.DoctorComment.E
	}
	body_text_list = []string{
		body_text,
	}
	// bubble_body_list := createBubbleBodyContents(body_text_list)
	// bubble_contents_item := createBubbleContentsItem("医師からのコメント", bubble_body_list)
	bubble_list = append(bubble_list, joinContents(header_text, body_text_list))

	// BMIの値からBMI Bubbleを生成
	header_text = "BMI値について"
	bmi_main_message := "あなたのBMI値は" + strconv.FormatFloat(answer_params.BMI, 'f', 2, 64) + "です。\n"
	switch {
	case answer_params.BMI <= 18.5: body_text = result_message.BMIComment.Low
	case answer_params.BMI >= 25: body_text = result_message.BMIComment.High
	default: body_text = " "
	}
	body_text_list = []string{
		bmi_main_message,
		result_message.BMIComment.Base,
		body_text,
	}
	bubble_list = append(bubble_list, joinContents(header_text, body_text_list))

	// 精液検査の回答結果から精液Bubbleを生成
	// 既婚者で相手が検査していない場合
	if answer_params.Married == "yes" {
		header_text = "精液検査について"
		// 相手が検査していない場合
		if answer_params.SemenExam == "no" {
			body_text_list = []string{result_message.SemenComment.Married}
			bubble_list = append(bubble_list, joinContents(header_text, body_text_list))
		} else if answer_params.SemenExam == "yes" {
			// 相手が検査していて、値が低い場合
			if answer_params.Semen == "20-39" || answer_params.Semen == "0-19" || answer_params.Semen == "unknown" {
				body_text_list = []string{result_message.SemenComment.MarriedAndLow}
				bubble_list = append(bubble_list, joinContents(header_text, body_text_list))
			}
		}
	}

	// AMHの回答結果から　AMH Bubbleを生成
	// AMH検査を受けていない場合
	if answer_params.AMHExam == "no" {
		header_text = "AMH検査について"
		body_text_list = []string{result_message.AMHComment.NotTaken}
		bubble_list = append(bubble_list, joinContents(header_text, body_text_list))
	} else if answer_params.AMHExam == "yes" {
		// 年齢が35歳以上でAMHの値が1.5未満
		age, _ := strconv.Atoi(answer_params.Age)
		amh, _ := strconv.ParseFloat(answer_params.AMH, 64)
		if age >= 35 && amh < 1.5 {
			header_text = "AMH検査について"
			body_text_list = []string{result_message.AMHComment.Low}
			bubble_list = append(bubble_list, joinContents(header_text, body_text_list))
		}
	}

	// 生理周期の回答結果から生理周期Bubbleを生成
	if answer_params.MenstrualCycle == "no" {
		header_text = "生理周期について"
		body_text_list = []string{result_message.MenstrualCycleComment.No}
		bubble_list = append(bubble_list, joinContents(header_text, body_text_list))
	}

	// Carouselを生成
	contents := createCarousel(bubble_list)
	fmt.Println(contents)

	return contents
}

func joinContents(header_text string, body_text_list []string) interface{} {
	bubble_body_list := createBubbleBodyContents(body_text_list)
	bubble_contents_item := createBubbleContentsItem(header_text, bubble_body_list)

	return bubble_contents_item
}

func createBubbleBodyContentsItem(body_text string) interface{} {
	bubble_body_contents_item := map[string]interface{}{
		"type":   "text",
		"text":   body_text,
		"color":  "#8C8C8C",
		"size":   "sm",
		// "weight": "bold",　
		"wrap":   true,
	}

	return bubble_body_contents_item
}

func createBubbleBodyContents(body_text_list []string) interface{} {
	bubble_body_contents_list := []interface{}{}
	for _, s := range body_text_list {
		bubble_body_contents_item := createBubbleBodyContentsItem(s)
		bubble_body_contents_list = append(bubble_body_contents_list, bubble_body_contents_item)
	}

	bubble_body_list := map[string]interface{}{
		"type":   "box",
		"layout": "vertical",
		"contents": []interface{}{
			map[string]interface{}{
				"type":     "box",
				"layout":   "vertical",
				"contents": bubble_body_contents_list,
				"flex":     1,
			},
		},
		"spacing":    "md",
		"paddingAll": "12px",
	}

	return bubble_body_list
}

func createBubbleContentsItem(header_text string, body_contents interface{}) interface{} {
	bubble_contents_item := map[string]interface{}{
		"type": "bubble",
		"size": "mega",
		"header": map[string]interface{}{
			"type": "box",
			"layout": "vertical",
			"contents": []interface{}{
				map[string]interface{}{
					"type": "text",
					"text": header_text,
					// "color": "#ffffff",
					"align": "start",
					"size": "md",
					"gravity": "center",
				},
			},
			"backgroundColor": "#FF6B6E",
			"paddingTop": "19px",
			"paddingAll": "12px",
			"paddingBottom": "16px",
		},
		"body": body_contents,
		"styles": map[string]interface{}{
			"footer": map[string]interface{}{
				"separator": false,
			},
		},
	}

	return bubble_contents_item
}

func createCarousel(bubble_contents_list []interface{}) string {
	map_result_contents := map[string]interface{}{
		"type": "carousel",
		"contents": bubble_contents_list,
	}
	json_result_contents, _ := json.Marshal(map_result_contents)
	result_contents := string(json_result_contents)
	return result_contents
}