package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"kedi_uz_bot/models"
	"log"
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)


func stats(b *gotgbot.Bot, ctx *ext.Context) error {
	res, err := http.Get("https://kedi.uz/api/v1/book/lost-animal-list/")
	if err != nil {
		return fmt.Errorf("failed to fetch data lost-animal-list/: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var lostAnimal []models.LostAnimal
	err = json.Unmarshal(body, &lostAnimal)
	if err != nil {
		log.Fatal(err)
	}

	msg := "<b>Lost Animals:</b>\n\n"
	for _, a := range lostAnimal {
		msg += fmt.Sprintf("🐾 <b>%s</b>\n📍 %s\n🗓️ %s\n\n", a.Title, a.Location, a.DateLost)
	}
	_, err = b.SendMessage(ctx.EffectiveUser.Id, msg, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	})

	if err != nil {
		return fmt.Errorf("failed to send /help handler message: %w", err)
		
	}
	return nil
}