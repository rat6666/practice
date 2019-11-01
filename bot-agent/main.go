package main

import (
	b "bot-agent/bot"
	"errors"
	"fmt"
)

func main() {
	bot, err := ManualBot("English")
	if err != nil {
		fmt.Println(err)
	}
	bot.SayHello()

	fmt.Println("Type english for english bot or / наберите русский(russian) для русского бота.")
	b.CreateBot(b.UserInput("Language/язык"))
}

func ManualBot(language string) (b.Sayer, error) {
	var bot b.Sayer
	var err error
	switch {
	case language == "English":
		bot = b.BotEnglish{Name: "John", Appeal: "You"}
	case language == "Russian":
		bot = b.BotRussian{Name: "Иван", Appeal: "Вы"}
	default:
		err = errors.New("Something wrong")
	}
	return bot, err
}
