package main

import (
	"encoding/json"
	"fmt"

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
	contents, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonString))
	if err != nil {
		return err
	}
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
		switch event.Type {
		case linebot.EventTypeMessage:
			event_json, _ := json.Marshal(&event)
			fmt.Println(string(event_json))

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// r.handleText(message, event.ReplyToken, event.Source.UserID)
				event_json, _ := json.Marshal(&event)
				fmt.Println(string(event_json), message)
				message_text := message.Text
				if message_text == "test" {
					fmt.Println("this is test message")
					r.handleFlex(Create_contents("test message"), event.ReplyToken, event.Source.UserID)
				}
				r.handleFlex(Create_contents("insert"), event.ReplyToken, event.Source.UserID)
			}
		}
	}
}

func (r *Line) handleText(message *linebot.TextMessage, replyToken, userID string) {
	r.SendTextMessage(message.Text, replyToken)
}

func (r *Line) handleFlex(contents string, replyToken, userID string) {
	altText := "this is alt text"
	r.SendFlexMessage(altText, contents, replyToken)
}
