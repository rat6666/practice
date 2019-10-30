package bot

import (
	"fmt"
	"strings"
)

type Bot struct {
	BotEnglish
	BotRussian
}

func Start() {
	fmt.Println("Input English for english bot or Введите Russian для русского бота.")
	language := ""
Begin:
	for {
		_, err := fmt.Scanln(&language)
		if err != nil {
			fmt.Print("Wrong input:", err, "\n")
			goto Begin
		}
		switch {
		case strings.Contains("engenglishутпутпдшыр", strings.ToLower(language)):
			fmt.Println("gotcha")
			var bot Sayer = BotEnglish{}
			bot.SayHello()
			bot.SayTime()
		}
		fmt.Println("Try again/ Попробуйте снова")
	}
}
