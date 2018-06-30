package main

import(
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strings"
	"strconv"
	"os"
	"net/http"
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
		log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + secret)
	go http.ListenAndServe("0.0.0.0:" + port, nil)

	for update := range updates {
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
