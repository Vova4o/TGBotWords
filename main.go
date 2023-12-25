package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func init() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Error loading .env file")
		os.Exit(1)
	}
}

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("С начала", "start"),
		tgbotapi.NewInlineKeyboardButtonData("Буква", "letter"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Цифра", "number"),
		tgbotapi.NewInlineKeyboardButtonData("Ввести букву", "enterLetter"),
	),
)

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)
	// Получаем обновления из канала updates
	// и обрабатываем каждое по очереди
	for update := range updates {

		if update.CallbackQuery != nil {
			callback := update.CallbackQuery
			callbackData := callback.Data

			if callbackData != "" {
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, callback.Data)
				msg.Text = "Вы выбрали " + callback.Data
				bot.Send(msg)
			}
		}

		// Проверяем, что сообщение не пустое
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		// конструируем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста напишите /start для начала работы с ботом")

		switch update.Message.Text {
		case "/start":
			msg.Text = "Привет, добро пожаловать в наш бот!\n Ниже представлены кнопки для навигации по боту."
			msg.ReplyMarkup = numericKeyboard
		case "start":
			msg.Text = "Привет, добро пожаловать в наш бот!\n Ниже представлены кнопки для навигации по боту."
			msg.ReplyMarkup = numericKeyboard
		}

		// и отправляем его обратно
		if _, err := bot.Send(msg); err != nil {
			log.Print(err)
		}
	}

}
