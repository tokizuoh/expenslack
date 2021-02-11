package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	slackToken := os.Getenv("SLACK_BOT_TOKEN")
	slackMemberId := os.Getenv("SLACK_MEMBER_ID")

	api := slack.New(slackToken)
	user, err := api.GetUserInfo(slackMemberId)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}
