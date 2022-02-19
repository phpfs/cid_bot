package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Panic("$PORT must be set!")
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Panic("$TOKEN must be set!")
	}

	url := os.Getenv("URL")
	if url == "" {
		log.Panic("$URL must be set!")
	}

	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Panic("$SECRET must be set!")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s!\n", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url + "/" + secret))
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("[Last Error %d] %s", info.LastErrorDate, info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + secret)
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+port, nil)
		if err != nil {
			log.Fatalln("[Webserver Error]", err)
		}
	}()

	for update := range updates {
		if update.Message == nil || update.Message.Chat == nil {
			continue
		}

		log.Println("--> Received message!")

		var msg tgbotapi.MessageConfig

		if strings.Contains(update.Message.Text, "/start") {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to @cid_bot!\n\nYour ChatID is: <b>"+strconv.Itoa(int(update.Message.Chat.ID))+"</b>\n\nIf you want to know a little more about this bot, send /about")
		} else if strings.Contains(update.Message.Text, "/about") {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "@cid_bot was built by phpfs and its source code is open sourced on github.com/phpfs/cid_bot. Currently, @cid_bot serves you from Heroku :)")
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Your ChatID is:\n<b>"+strconv.Itoa(int(update.Message.Chat.ID))+"</b>\n\nChatIDs normally don't change, but you can ask me at any time with /chatid what your current ChatID is :)")
		}

		msg.ParseMode = "HTML"
		_, err := bot.Send(msg)

		if err != nil {
			log.Println("There was error sending the last message!", err)
		}
	}
}
