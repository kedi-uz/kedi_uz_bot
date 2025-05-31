package handlers

import (
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

const (
	DISTRICT = "district"
)

func Dispatcher() *ext.Dispatcher{
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error)  ext.DispatcherAction{
			log.Println("an error occurred while handling update:", err.Error())
			
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})

	// dispatcher.AddHandler(handlers.NewCommand("start", start))
	dispatcher.AddHandler(handlers.NewConversation(
		[]ext.Handler{handlers.NewCommand("start", start)},
		map[string][]ext.Handler{
			DISTRICT: {handlers.NewMessage(nil, District)},
		},
		&handlers.ConversationOpts{
			// Exits: []ext.Handler{handlers.NewCommand("cancel", Cancel)},
			// StateStorage: conversation.NewInMemoryStorage(conversation.KeyStrategySenderAndChat),
			// AllowReEntry: true,
		},
	))
	dispatcher.AddHandler(handlers.NewCommand("about", about))
	dispatcher.AddHandler(handlers.NewCommand("help", help))
	// dispatcher.AddHandler(handlers.NewCommand("stats", stats))
	
	return dispatcher
}