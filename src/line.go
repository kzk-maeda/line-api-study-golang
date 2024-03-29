package main

import (
	"encoding/json"
	"fmt"

	"router"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Line struct {
	ChannelSecret string
	ChannelToken  string
	Bot           *linebot.Client
}

func (r *Line) SendTextMessage(message string, replyToken string) error {
	return r.Reply(replyToken, linebot.NewTextMessage(message))
}

func (r *Line) SendFlexMessage(altText string, jsonString string, replyToken string) error {
	fmt.Println(jsonString)
	contents, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonString))
	if err != nil {
		fmt.Println("FlexMessage Syntax Error")
		fmt.Println(err)
		return err
	}
	fmt.Println("Send FlexMessage as Below.")
	// fmt.Println(string(contents))
	return r.Reply(replyToken, linebot.NewFlexMessage(altText, contents))
}

func (r *Line) Reply(replyToken string, message linebot.SendingMessage) error {
	if _, err := r.Bot.ReplyMessage(replyToken, message).Do(); err != nil {
		fmt.Printf("Reply Error: %v", err)
		return err
	}
	return nil
}

func (l *Line) New(secret, token string) error {
	l.ChannelSecret = secret
	l.ChannelToken = token

	bot, err := linebot.New(
		l.ChannelSecret,
		l.ChannelToken,
	)
	if err != nil {
		return err
	}

	l.Bot = bot
	return nil
}

func (r *Line) EventRouter(eve []*linebot.Event) {
	for _, event := range eve {
		event_json, _ := json.Marshal(&event)
		user_id := event.Source.UserID
		reply_token := event.ReplyToken
		replyContents := ""
		fmt.Println(string(event_json))

		switch event.Type {
		case linebot.EventTypeMessage: // MessageEventを受け取ったとき（初回と値入力）

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				fmt.Println(string(event_json), message)
				replyContents = router.MessageRouter(user_id, message.Text)
			}
			// TODO: error handling

		case linebot.EventTypePostback: // PostbackEventを受け取ったとき（選択肢に対する回答）
			eventData := event.Postback.Data
			replyContents = router.PostbackRouter(user_id, eventData)
		}
		r.handleFlex(replyContents, reply_token, user_id)
	}
}

func (r *Line) handleText(message *linebot.TextMessage, replyToken, userID string) {
	r.SendTextMessage(message.Text, replyToken)
}

func (r *Line) handleFlex(contents, replyToken, userID string) {
	altText := "診断ツール"
	r.SendFlexMessage(altText, contents, replyToken)
}
