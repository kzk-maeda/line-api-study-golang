package dto

type DoctorComment struct {
	A string `json:"A"`
	B string `json:"B"`
	C string `json:"C"`
	D string `json:"D"`
	E string `json:"E"`
}

type BMIComment struct {
	Base string `json:"base"`
	Low string `json:"low"`
	High string `json:"high"`
}

type SemenComment struct {
	Married string `json:"married"`
	MarriedAndLow string `json:"married_and_low"`
}

type AMHComment struct {
	NotTaken string `json:"not_taken"`
	Low string `json:"low"`
}

type MenstrualCycleComment struct {
	No string `json:"no"`
}

type ResultMessage struct {
	DoctorComment DoctorComment `json:"doctor_comment"`
	BMIComment BMIComment `json:"bmi_comment"`
	SemenComment SemenComment `json:"semen_comment"`
	AMHComment AMHComment `json:"amh_comment"`
	MenstrualCycleComment MenstrualCycleComment `json:"menstrual_cycle_comment"`
}