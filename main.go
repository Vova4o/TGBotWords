package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var Arr []string
var StringNewArr string

func init() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Error loading .env file")
		os.Exit(1)
	}
}

func init() {
	Arr, _ = readTextFile()
}

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

		// fmt.Println(len(Arr))
		// получает callback от кнопок
		if update.CallbackQuery != nil {
			callback := update.CallbackQuery
			callbackData := callback.Data

			// работаем с callbackData, отвечаем на запросы кнопок
			switch callbackData {
			case "fpage":
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, callback.Data)
				if len(Arr) < 100 {
					msg.Text = strings.Join(Arr, ", ")
				} else {
					msg.Text = strings.Join(Arr[:100], ", ")
				}
				msg.ReplyMarkup = numericKeyboardMidl
				bot.Send(msg)
			case "lpage":
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, callback.Data)
				if len(Arr) < 100 {
					msg.Text = strings.Join(Arr, ", ")
				} else {
					msg.Text = strings.Join(Arr[len(Arr)-100:], ", ")
				}
				msg.ReplyMarkup = numericKeyboardLast
				bot.Send(msg)
			case "back":
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, callback.Data)
				// need to fix it
				if len(Arr) < 100 {
					msg.Text = strings.Join(Arr, ", ")
				} else {
					msg.Text = strings.Join(Arr[len(Arr)-100:], ", ")
				}
				msg.ReplyMarkup = numericKeyboardMidl
				bot.Send(msg)
			case "forward":
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, callback.Data)
				// need to fix it
				if len(Arr) < 100 {
					msg.Text = strings.Join(Arr, ", ")
				} else {
					msg.Text = strings.Join(Arr[:100], ", ")
				}
				msg.ReplyMarkup = numericKeyboardMidl
				bot.Send(msg)
			}
		}

		// Проверяем, что сообщение не пустое
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		// конструируем ответное сообщение
		// ЭТО ДЕФОЛТНОЕ СООБЩЕНИЕ, ДАЛЕЕ НАЧИНАЕМ С НИМ ИГРАТЬ
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Если вы любите играть в слова,\nто данный бот поможет вам найти слова по буквам\nПожалуйста, нажмите *start* для начала работы с ботом")
		// делаем шрифт жирным
		msg.ParseMode = "markdown"
		msg.ReplyMarkup = mainKeyboard

		numOfLetters, err := strconv.Atoi(update.Message.Text)
		if err != nil {
			log.Print(err)
		} else {
			Arr = shrinkByLen(Arr, numOfLetters)
			// StringNewArr = strings.Join(Arr, ", ")
			var StringArr string
			if len(Arr) < 100 {
				StringArr = strings.Join(Arr, ", ")
			} else {
				StringArr = strings.Join(Arr[:100], ", ")
			}
			msg.Text = StringArr + "\n\nЕсли вы хотите ограничить колличество букв в слове, то введите цифру."
			msg.ReplyMarkup = numericKeyboardMidl
		}

		// путем сложных манипуляций с текстом рунами и еще чем попало
		// мы получаем первую букву введенного слова
		str := strings.ToLower(update.Message.Text)
		firstChar := str[0]
		a := "а"
		z := "я"
		aByte := a[0]
		zByte := z[0]
		if firstChar >= aByte && firstChar <= zByte {
			fmt.Println(str)
			Arr = findMatch(Arr, str)
			// StringNewArr = strings.Join(Arr, ", ")
			var StringArr string
			if len(Arr) < 100 {
				StringArr = strings.Join(Arr, ", ")
			} else {
				StringArr = strings.Join(Arr[:100], ", ")
			}
			msg.Text = StringArr + "\n\nТелеграм позволят показывать 4096 символов в сообщении, продолжайте выборку\nЕсли вы хотите ограничить колличество букв в слове, то введите цифру."
			msg.ReplyMarkup = numericKeyboardMidl
		}

		if firstChar >= byte('a') && firstChar <= byte('z') {
			msg.Text = "Вы ввели букву " + update.Message.Text + ", я пока маленький и не понимаю Ангийский язык, но я учусь, попробуйте ввести русскую букву"
		}

		switch update.Message.Text {
		case "/start":
			msg.Text = "Привет, добро пожаловать в словестный бот!\nОтправьте боту букву и он выдаст вам список слов с этой бувой, если вы хотите ограничить колличество букв в слове, то введите цифру.\nДля того чтобы начать с начала выберете команду *reset*."
			// msg.ReplyMarkup = numericKeyboard
			msg.ParseMode = "markdown"
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		case "/reset":
			msg.Text = "Вы выбрали команду *reset* начинайте выбор слов с начала."
			Arr = Reset() // ресетим слова в боте, тупо перезаписываем оригинальный массив
			msg.ParseMode = "markdown"
		}

		// и отправляем его обратно
		if _, err := bot.Send(msg); err != nil {
			log.Print(err)
		}
	}

}
