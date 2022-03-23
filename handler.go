package main

import "fmt"

func handleRedeem(bc *BotCommand) {
	bc.client.Say(bc.message.Channel, fmt.Sprintf("%s resgatou 1 token", bc.message.User.DisplayName))
}

func handleWallet(bc *BotCommand) {
	fmt.Println("handleWallet")
}

func handleBalance(bc *BotCommand) {
	fmt.Println("handleBalance")
}
