package main

import (
	"log"
	"os"

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

	for update := range updates {
		if update.Message != nil { // если получаем не пустое сообщение
			switch update.Message.Command() { // получаем команду из сообщения
			case "help", "start": // помощь, справочная информация
				helpCommand(bot, update.Message)
			case "list":
				listCommand(bot, update.Message, productService)
			default: // сообщение без команды
				defaultBehaviour(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products",
	)

	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService *product.Service) {
	outputMsgText := "Here are all products: \n\n"

	products := productService.List()
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	bot.Send(msg)
}

func defaultBehaviour(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text) // логируем сообщение

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text) // создаем ответ, указываем id чата и текст сообщения
	//msg.ReplyToMessageID = update.Message.MessageID                                       // указание того, что сообщение - это реплай на предыдущее сообщение

	bot.Send(msg)
}
