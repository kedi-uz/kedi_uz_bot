package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"kedi_uz_bot/configs"
	"kedi_uz_bot/handlers"
	"kedi_uz_bot/models"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func main() {

	conf := configs.LoadConfig()

	// db connection
	configs.InitDB()

	// Get token from the environment variable
	token := conf.TelegramApiToken
	if token == "" {
		panic("TOKEN environment variable is invalid")
	}

	// Create bot from environment value.
	b, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		BotClient: &gotgbot.BaseBotClient{
			Client: http.Client{},
			DefaultRequestOpts: &gotgbot.RequestOpts{
				Timeout: gotgbot.DefaultTimeout, // Customise the default request timeout here
				APIURL:  gotgbot.DefaultAPIURL,  // As well as the Default API URL here (in case of using local bot API servers)
			},
		},
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	dispatcher := handlers.Dispatcher()
	updater := ext.NewUpdater(dispatcher, nil)

	// Start receiving updates.
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...\n", b.User.Username)

	
	// Post request listener in this endpoint
	server := &Server{Bot: b}
	http.HandleFunc("/notify-lost-animal", server.handleLostAnimalNotification)
	log.Println("HTTP server running on :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()

}

type Server struct {
		Bot *gotgbot.Bot
	}

func (s *Server) handleLostAnimalNotification(w http.ResponseWriter, r *http.Request) {
	var lostAnimal models.LostAnimal
	err := json.NewDecoder(r.Body).Decode(&lostAnimal)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	msg := fmt.Sprintf("<b>%s tumanida hayvon yoqoldi</b> \n\n", lostAnimal.District.Title)
	msg += fmt.Sprintf("🐾 <b>%s</b>\n📍 %s\n🗓️ %s\n\n", lostAnimal.Title, lostAnimal.Location, lostAnimal.DateLost)
	
	var users []models.TelegramUser
	db := configs.DB

	result := db.Find(&users)
	if result.Error != nil {
		log.Printf("failed to retrieve users: %v", result.Error)
		return
		}

	for _, user := range users {
		if user.District == lostAnimal.District.Title {
			_ , err = s.Bot.SendMessage(user.TelegramID, msg, &gotgbot.SendMessageOpts{
				ParseMode: "HTML",})

			if err != nil {
				log.Printf("Failed to send Telegram message: %v", err)
				http.Error(w, "Failed to send Telegram message", http.StatusInternalServerError)
				return
			}
		}
	}

	w.Write([]byte("Notifications sent"))
}