package router

import (
	"fmt"
	"math"
	"strconv"
	"router/repository"
)

func calculateRank(user_id string) string {
	// Databaseから対象user_idの結果を取得
	data, _ := repository.GetData(user_id)
	fmt.Println(data.Answers)

	// 結果の各値から診断結果を計算
	score := 0
	// Question 2-1
	if val, ok := data.Answers["2-1"]; ok {
		switch val.Answer {
		case "1":
			score += 0
		case "2":
			score += 3
		case "3-4":
			score += 7
		case "5-6":
			score += 12
		case "7":
			score += 18
		}
	}

	// Question 2-2
	if val, ok := data.Answers["2-2"]; ok {
		switch val.Answer {
		case "yes":
			score += 4
		case "no":
			score += 0
		}
	}

	// Question 2-4
	if val, ok := data.Answers["2-4"]; ok {
		switch val.Answer {
		case "60":
			score += 0
		case "40-59":
			score += 2
		case "20-39":
			score += 4
		case "0-19":
			score +=6
		case "unknown":
			score += 3
		}
	}

	// Question 3-1
	if val, ok := data.Answers["3-1"]; ok {
		age, _ := strconv.Atoi(val.Answer)
		if age <= 25 {
			score += 0
		} else if 26 <= age && age <= 31 {
			score += 3
		} else if 32 <= age && age <= 35 {
			score += 7
		} else if 36 <= age && age <= 37 {
			score += 10
		} else if 38 <= age && age <= 39 {
			score += 11
		} else if 40 <= age {
			score += 15
		}
	}

	// Question 6-1
	if val, ok := data.Answers["6-1"]; ok {
		switch val.Answer {
		case "yes":
			score += 0
		case "no":
			score += 6
		}
	}

	// Question 6-6
	if val, ok := data.Answers["6-6"]; ok {
		switch val.Answer {
		case "level_1":
			score += 6
		case "level_2":
			score += 8
		case "level_3":
			score += 12
		case "level_0":
			score += 0
		}
	}

	fmt.Println("Score: ", score)
	
	// result
	rank := ""
	if score <= 9 {
		rank = "A"
	} else if 10 <= score && score <= 19 {
		rank = "B"
	} else if 20 <= score && score <= 34 {
		rank = "C"
	} else if 35 <= score && score <= 49 {
		rank = "D"
	} else if 50 <= score {
		rank = "E"
	}

	return rank
}

func calculateBMI(user_id string) float64 {
	// Databaseから対象user_idの結果を取得
	data, _ := repository.GetData(user_id)

	// BMIを計算
	height, _ := strconv.Atoi(data.Answers["3-2"].Answer)
	weight, _ := strconv.Atoi(data.Answers["3-3"].Answer)
	fmt.Println(height, " ", weight)

	raw_BMI := float64(weight) / (float64(height)/100 * float64(height)/100)
	// math.Round(v*100) / 100
	fmt.Println(raw_BMI)
	BMI := math.Round(raw_BMI * 10) / 10
	fmt.Println("BMI : ", BMI)

	return BMI
}