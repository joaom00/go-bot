package main

type Handler func(*BotCommand)

type FindHandler func(string) (Handler, bool)

type BotCommand struct {
	findHandler FindHandler
}

type Router struct {
	rules map[string]Handler
}

func NewBotCommand(findHandler FindHandler) *BotCommand {
	return &BotCommand{
		findHandler: findHandler,
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
