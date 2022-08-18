package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"time"
)

var duration = time.Minute

var m time.Time

func newBot() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	os.Setenv("token", halyava007BotToken)
	bot, err := tgbotapi.NewBotAPI(os.Getenv("token"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	return bot, updates
}
func main() {
	//конектимся к телеграм и получаем апдейты
	bot, updates := newBot()
	//создаём кеш
	c := New()
	//иницализируем структуру с bot и кешем
	x := newXrr(bot, c)
	//загружаем в кеш Юзеров из Excel
	loadFromExcel(c)
	// повторяющиеся сообщения
	if err := callAt(21, 00, 0, x.sendMessageAtTime, 24*time.Hour); err != nil {
		fmt.Printf("error while 'callAt': %v\n", err)
	}
	// запись в excel всех кто оплатил билеты, запись n часов
	if err := callAt(20, 50, 0, c.writeInFile, 6*time.Hour); err != nil {
		fmt.Printf("error while 'callAt': %v\n", err)
	}

	//запись кэша в excel файл
	callAtStartOfProgram(c.writeCacheInFile, 5*duration)
	//фейковое пополнение банка
	callAtStartOfProgram(bankAddFakeMoney, 15*duration)
	// Loop through each update.
	getUpdates(updates, bot, c)

}
