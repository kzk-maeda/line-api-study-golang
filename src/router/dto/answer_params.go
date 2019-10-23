package dto

type AnswerParams struct {
	Rank string `json:"rank"`
	BMI float64 `json:"bmi"`
	Married string `json:"married"`
	SemenExam string `json:"semen_exam,omitempty"`
	Semen string `json:"semen,omitempty"`
	Age string `json:"age"`
	AMHExam string `json:"amh_exam,omitempty"`
	AMH string `json:"amh,omitempty"`
	MenstrualCycle string `json:"menstrual_cycle"`
}