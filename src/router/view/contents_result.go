package view

import (
	"fmt"
	"encoding/json"
	"router/dto"
)

func CreateResult(answer_params dto.AnswerParams) string {
	// message定義を読み込み
	result_message := resultMessage()
	
	bubble_list := []interface{}{}

	// rankの値から診断結果Bubbleを生成
	main_result_text := "あなたの妊娠力はランク" + answer_params.Rank + "です"
	main_body_text_list := []string{
		main_result_text,
		"詳細は右にスワイプしてください",
	}
	main_bubble_body_list := createBubbleBodyContents(main_body_text_list)
	main_bubble_contents_item := createBubbleContentsItem("診断結果：ランク" + answer_params.Rank, main_bubble_body_list)
	bubble_list = append(bubble_list, main_bubble_contents_item)

	// rankの値から医師コメントBubbleを生成
	doctor_body_text := ""
	switch answer_params.Rank {
		case "A": doctor_body_text = result_message.DoctorComment.A
		case "B": doctor_body_text = result_message.DoctorComment.B
		case "C": doctor_body_text = result_message.DoctorComment.C
		case "D": doctor_body_text = result_message.DoctorComment.D
		case "E": doctor_body_text = result_message.DoctorComment.E
	}
	doctor_body_text_list := []string{
		doctor_body_text,
	}
	doctor_bubble_body_list := createBubbleBodyContents(doctor_body_text_list)
	doctor_bubble_contents_item := createBubbleContentsItem("医師からのコメント", doctor_bubble_body_list)
	bubble_list = append(bubble_list, doctor_bubble_contents_item)

	// BMIの値からBMI Bubbleを生成

	// 精液検査の回答結果から精液Bubbleを生成

	// AMHの回答結果から　AMH Bubbleを生成

	// 生理周期の回答結果から生理周期Bubbleを生成

	// Carouselを生成
	contents := createCarousel(bubble_list)
	fmt.Println(contents)

	return contents
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