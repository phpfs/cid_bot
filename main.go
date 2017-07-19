package main

import(
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strings"
	"strconv"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("Token")
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
  	if err != nil {
    		log.Panic(err)
  	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if(strings.Contains(update.Message.Text, "/start")){
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to @cid_bot!\n\nStart by querying your Telegram Chat_ID:\n/chatid\n\nIf you want to know something about this bot, send:\n/about\n\nGreetings, phpfs")
			bot.Send(msg)
		}else if(strings.Contains(update.Message.Text, "/about")){
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "@cid_bot was built by phpfs and its source code is open sourced on github.com/phpfs/cid_bot!")
			bot.Send(msg)
		}else{
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Your ChatID is:\n\n" + strconv.Itoa(int(update.Message.Chat.ID)))
			bot.Send(msg)
		}
	}
}
