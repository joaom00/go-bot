package main

import (
	"fmt"
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

	commands := map[string]string{"!resgatar": "ResgatarHandler", "!carteira": "CarteiraHandler", "!saldo": "SaldoHandler"}

	client := twitch.NewClient(twitchUsename, twitchOauthToken)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
		isCommand := strings.HasPrefix(message.Message, "!")
		if isCommand {
			commandName := strings.Split(message.Message, " ")

			test := commands[commandName[0]]

			client.Say(twitchChannel, test)
		}
	})

	client.Join(twitchChannel)

	err = client.Connect()
	if err != nil {
		panic(err)
	}
}
