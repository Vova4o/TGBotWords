package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Error loading .env file")
		os.Exit(1)
	}
}

//var TgBotToken = os.Getenv("TG_BOT_TOKEN")
//var BotApi = os.Getenv("BOT_API")

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
		// Проверяем, что сообщение не пустое
		if update.Message == nil { // ignore non-Message updates
			continue
		}
		//// Создаем текст сообщения для отправки пользователю
		//text := "Hello Welcome to our bot!"
		//// Создаем сообщение для отправки пользователю
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// Используем switch для выбора действия по тексту сообщения
		switch update.Message.Text {
		case "/start":
			msg.Text = "Hello Welcome to our bot! Swithc!"
		}
		// Отправляем сообщение пользователю
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

}
