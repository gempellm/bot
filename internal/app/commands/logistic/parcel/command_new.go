package parcel

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gempellm/bot/internal/model/logistic"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) New(inputMsg *tgbotapi.Message) {
	rawArgs := inputMsg.CommandArguments()
	args := strings.Split(rawArgs, " ")

	mu := &sync.Mutex{}

	for _, arg := range args {
		if len(arg) > 0 {
			go c.saveArg(arg, inputMsg, mu)
		}
	}

}

func (c *Commander) saveArg(title string, inputMsg *tgbotapi.Message, mu *sync.Mutex) {
	mu.Lock()
	var lastID uint64
	data, _ := os.ReadFile("lastID.txt")
	lastID, _ = strconv.ParseUint(string(data), 10, 64)
	lastID++
	os.WriteFile("lastID.txt", []byte(fmt.Sprint(lastID)), 0666)

	parcel := logistic.Parcel{Title: title, ParcelID: lastID, Timestamp: time.Now().Unix()}

	parcelID, err := c.parcelService.Create(parcel)
	mu.Unlock()
	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Error occured during parcel creation.")
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Successfully created parcel:\nTitle: %s\nparcelID: %d", title, parcelID))
	c.bot.Send(msg)
}
