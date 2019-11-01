package bot

import (
	"fmt"
)

func CreateBot(language string) {
	for {
		switch {
		case language == "english":
			var bot = BotEnglish{Name: "John", Appeal: "You"}
			Start(bot, bot.Appeal)
		case language == "russian" || language == "русский":
			var bot = BotRussian{Name: "Иван", Appeal: "Вы"}
			Start(bot, bot.Appeal)
		default:
			fmt.Println("Try again")
			CreateBot(UserInput("Language/язык"))
		}
	}
}

func Start(bot Sayer, appeal string) {
	bot.SayHello()
	for {
		input := UserInput(appeal)
		switch {
		case input == "1" || input == "hello" || input == "привет":
			bot.SayHello()
		case input == "2" || input == "time" || input == "время":
			bot.SayTime()
		case input == "3" || input == "date" || input == "дата":
			bot.SayDate()
		case input == "4" || input == "day" || input == "день":
			bot.SayDay()
		case input == "5" || input == "bye" || input == "пока":
			bot.SayBye()
		default:
			fmt.Println("Try again/Попробуйте снова")
		}
	}
}
