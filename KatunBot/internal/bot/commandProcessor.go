package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *Bot) commandProcessor(message *tgbotapi.Message) {
	command := message.Command()
	arg:= message.CommandArguments()

	switch command {
		case "vkat":
			b.regVkat(int64(message.From.ID), message.From.UserName, arg, message.MessageID, message.Chat.ID)

		case "deadline":
			b.deadline(int64(message.From.ID), message.MessageID,message.Chat.ID)

		case "katuni":
			b.vkatuns(message.MessageID,message.Chat.ID)

		case "untilvkat":
			b.daysUntilVkat(int64(message.From.ID), message.MessageID, message.Chat.ID)

		case "delete":
			b.DeleteVkat(int64(message.From.ID), arg, message.MessageID,message.Chat.ID)

		case "updatevkat":
			b.updateVkat(int64(message.From.ID), arg, message.MessageID, message.Chat.ID)
	}
}