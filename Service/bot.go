package Service

import (
	"log"

	botApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type CmdHandler func(string) string

type Bot struct {
	bot   *botApi.BotAPI
	route map[string]CmdHandler
}

var numericKeyboard = botApi.NewInlineKeyboardMarkup(
	botApi.NewInlineKeyboardRow(
		botApi.NewInlineKeyboardButtonData("ADD", "add"),
		botApi.NewInlineKeyboardButtonData("UPDATE", "update"),
		botApi.NewInlineKeyboardButtonData("DELETE", "delete"),
		botApi.NewInlineKeyboardButtonData("HELP", "help"),
		botApi.NewInlineKeyboardButtonData("LIST", "list"),
	),
)

func Init(apiKey string) (*Bot, error) {
	bot, err := botApi.NewBotAPI(apiKey)
	if err != nil {
		return nil, errors.Wrap(err, "init tgbot")
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Bot{
		bot:   bot,
		route: make(map[string]CmdHandler),
	}, nil
}

func (c *Bot) RegisterHandler(cmd string, f CmdHandler) {
	c.route[cmd] = f
}

func (c *Bot) Run() error {
	u := botApi.NewUpdate(0)
	u.Timeout = 60
	updates := c.bot.GetUpdatesChan(u)

	cmd := ""

	for update := range updates {
		if update.CallbackQuery != nil {
			if update.CallbackQuery.Data == "add" {
				cmd = "add"
				msg := botApi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Enter: name and age")
				c.bot.Send(msg)
			} else if update.CallbackQuery.Data == "update" {
				cmd = "update"
				msg := botApi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Enter: id,new name and new age")
				c.bot.Send(msg)
			} else if update.CallbackQuery.Data == "delete" {
				cmd = "delete"
				msg := botApi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Enter: id")
				c.bot.Send(msg)
			} else if update.CallbackQuery.Data == "help" {
				msg := botApi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
				f, _ := c.route["help"]
				msg.Text = f(update.CallbackQuery.Data)
				c.bot.Send(msg)
			} else if update.CallbackQuery.Data == "list" {
				msg := botApi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
				f, _ := c.route["list"]
				msg.Text = f(update.CallbackQuery.Data)
				c.bot.Send(msg)
			}

			continue
		}

		if update.Message == nil {
			continue
		}

		msg := botApi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.ReplyMarkup = botApi.NewRemoveKeyboard(true)
		}

		if cmd != "" {
			if f, ok := c.route[cmd]; ok {
				msg.Text = f(update.Message.Text)
			} else {
				msg.Text = "Unknown command"
			}
			cmd = ""
		} else {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg.Text = "Menu"
		}

		_, err := c.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send tg message")
		}
	}
	return nil
}
