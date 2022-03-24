package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/joho/godotenv"
)

func setInterval(p interface{}, interval time.Duration) chan<- bool {
	ticker := time.NewTicker(interval)
	stopIt := make(chan bool)

	go func() {
		for {
			select {
			case <-stopIt:
				fmt.Print("stop setInterval")
			case <-ticker.C:
				reflect.ValueOf(p).Call([]reflect.Value{})
			}
		}
	}()

	return stopIt
}

func handler(client *twitch.Client) {
	client.Say("brunobandev", "brdevstreamers")
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	twitchUsename := os.Getenv("TWITCH_USERNAME")
	twitchOauthToken := os.Getenv("TWITCH_OAUTH_TOKEN")
	twitchChannel := os.Getenv("TWITCH_CHANNEL")

	client := twitch.NewClient(twitchUsename, twitchOauthToken)

	stopper := setInterval(func() { handler(client) }, 2*time.Second)
	fmt.Println("your go program ...")

	// avoid the unused variable warn:
	_ = stopper
	// to stop setInterval uncomment the next line:
	// stopper <- true

	// pause the console
	<-make(chan bool)

	// client.Say(twitchChannel, commandName[0])

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {

		router := NewRouter()
		router.Handle("!resgatar", handleRedeem)
		router.Handle("!carteira", handleWallet)
		router.Handle("!saldo", handleBalance)

		bc := NewBotCommand(router.FindHandler, client, &message)
		command := strings.Split(message.Message, " ")[0]

		if handler, found := bc.findHandler(command); found {
			handler(bc)
		}
	})

	client.Join(twitchChannel)

	err = client.Connect()
	if err != nil {
		panic(err)
	}
}
