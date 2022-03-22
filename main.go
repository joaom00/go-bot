package main

import (
	"log"
	"os"
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	twitchUsename := os.Getenv("TWITCH_USERNAME")
	twitchOauthToken := os.Getenv("TWITCH_OAUTH_TOKEN")
	twitchChannel := os.Getenv("TWITCH_CHANNEL")

	client := twitch.NewClient(twitchUsename, twitchOauthToken)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {

		router := NewRouter()
		router.Handle("!resgatar", handleRedeem)
		router.Handle("!carteira", handleWallet)
		router.Handle("!saldo", handleBalance)

		bc := NewBotCommand(router.FindHandler)
		command := strings.Split(message.Message, " ")[0]

		if handler, found := bc.findHandler(command); found {
			handler(bc)
		}
		// client.Say(twitchChannel, commandName[0])

	})

	client.Join(twitchChannel)

	err = client.Connect()
	if err != nil {
		panic(err)
	}
}
