package main

import (
	"encoding/json"
)

// JSONのコンテンツを生成する関数
func Create_contents(h_text string) string {
	map_flex_contents := map[string]interface{}{
		"type": "bubble",
		"styles": map[string]interface{}{
			"header": map[string]interface{}{
				"backgroundColor": "#ffaaaa",
			},
			"body": map[string]interface{}{
				"backgroundColor": "#aaffaa",
			},
			"footer": map[string]interface{}{
				"backgroundColor": "#aaaaff",
			},
		},
		"header": map[string]interface{}{
			"type":   "box",
			"layout": "vertical",
			"contents": []interface{}{
				map[string]interface{}{
					"type": "text",
					"text": h_text,
				},
			},
		},
		"hero": map[string]interface{}{
			"type":        "image",
			"url":         "https://example.com/flex/images/image.jpg",
			"size":        "full",
			"aspectRatio": "2:1",
		},
		"body": map[string]interface{}{
			"type":   "box",
			"layout": "vertical",
			"contents": []interface{}{
				map[string]interface{}{
					"type": "text",
					"text": "body",
				},
				map[string]interface{}{
					"type": "text",
					"text": "body2",
				},
			},
		},
		"footer": map[string]interface{}{
			"type":   "box",
			"layout": "vertical",
			"contents": []interface{}{
				map[string]interface{}{
					"type": "text",
					"text": "footer",
				},
			},
		},
	}
	json_flex_contents, _ := json.Marshal(map_flex_contents)
	flex_contents := string(json_flex_contents)

	return flex_contents

}
