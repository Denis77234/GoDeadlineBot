package bot

import (
	"katun/internal/database/interface"
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct{
	Api *tgbotapi.BotAPI
	Database dbinterface.Database
}

func New(token string, db dbinterface.Database) (*Bot){
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	

	return &Bot{
		Api: bot,
		Database: db,
	}
}

func (b *Bot) sendMsg(text string, chatid int64, messageId int){
	msg:= tgbotapi.NewMessage(chatid, text)
	msg.ReplyToMessageID = messageId
	b.Api.Send(msg)
}

func (b *Bot) Start(){

	log.Printf("Authorized on account %s", b.Api.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates,err := b.Api.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {

		if update.Message != nil { 
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}

		b.commandProcessor(update.Message)


	}
}







