package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

var mainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/start"),
		tgbotapi.NewKeyboardButton("/reset"),
	),
)
