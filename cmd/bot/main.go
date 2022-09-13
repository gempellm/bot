package main

import (
	"log"
	"os"

	"github.com/gempellm/bot/internal/app/commands"
	"github.com/gempellm/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // подтягиваем переменные окружения из .env

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token) // создаем объект бота
	if err != nil {
		log.Panic(err) // если не вышло - падаем
	}

	bot.Debug = true // включаем режим дебага, который выводит все сообщения, которые пришли к боту

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	} // инициализируются настройки подписки на сообщения

	updates := bot.GetUpdatesChan(u) // получаем канал (очередь) сообщений

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		commander.HandleUpdate(update)
	}
}
