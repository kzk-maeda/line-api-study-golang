package view

func CreateResult(next_question string, rank string) string {
	question_text := "診断結果"
	textComponent := "あなたの妊娠力はランク" + rank + "です"
	bodyContents := createTextContents(textComponent)
	contents := createBaseContents(question_text, bodyContents)

	return contents
}