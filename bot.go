package main

import "github.com/gempir/go-twitch-irc/v3"

type Handler func(*BotCommand)

type FindHandler func(string) (Handler, bool)

type BotCommand struct {
	findHandler FindHandler
	client      *twitch.Client
	message     *twitch.PrivateMessage
}

type Router struct {
	rules map[string]Handler
}

func NewBotCommand(findHandler FindHandler, c *twitch.Client, m *twitch.PrivateMessage) *BotCommand {
	return &BotCommand{
		findHandler: findHandler,
		client:      c,
		message:     m,
	}
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

func (r *Router) Handle(command string, handler Handler) {
	r.rules[command] = handler
}

func (r *Router) FindHandler(command string) (Handler, bool) {
	handler, found := r.rules[command]
	return handler, found
}
