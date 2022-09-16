package parcel

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Delete(inputMsg *tgbotapi.Message) {
	rawData := inputMsg.CommandArguments()

	args := strings.Split(rawData, " ")

	mu := &sync.Mutex{}

	for _, arg := range args {
		go c.removeArg(arg, inputMsg, mu)
	}
}

func (c *Commander) removeArg(arg string, inputMsg *tgbotapi.Message, mu *sync.Mutex) {
	id, err := strconv.ParseUint(arg, 10, 64)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Wrong input format.")
		c.bot.Send(msg)
		return
	}

	mu.Lock()
	ok, _ := c.parcelService.Remove(id)
	mu.Unlock()

	if !ok {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Parcel with ID %d not found.", id))
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Parcel with ID %d was successfully removed.", id))
	c.bot.Send(msg)
}
