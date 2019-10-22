package dto

type AnswerParams struct {
	Rank string `json:"rank"`
	BMI float64 `json:"bmi"`
	Semen string `json:"semen,omitempty"`
	AMH string `json:"amh,omitempty"`
	MenstrualCycle string `json:"menstrual_cycle"`
}