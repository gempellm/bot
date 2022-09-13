package main

import (
	"log"
	"os"

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

	for update := range updates {
		if update.Message != nil { // если получаем не пустое сообщение
			switch update.Message.Command() { // получаем команду из сообщения
			case "help": // /help
				helpCommand(bot, update.Message)
			default: // сообщение без команды
				defaultBehaviour(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help")

	bot.Send(msg)
}

func defaultBehaviour(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text) // логируем сообщение

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text) // создаем ответ, указываем id чата и текст сообщения
	//msg.ReplyToMessageID = update.Message.MessageID                                       // указание того, что сообщение - это реплай на предыдущее сообщение

	bot.Send(msg)
}
