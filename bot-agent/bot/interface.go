package bot

import (
	"fmt"
)

type Sayer interface {
	SayHello()
	SayTime()
	// SayDate()
	// SayDay()
	// SayBye()
}

func (b Bot) SayHello() {
	fmt.Println("Hello, I am John ")
}
func (b Bot) SayTime() {
	fmt.Println("Hello, I am John ")
}

// func (b BotEnglish) SayDate() {
// 	fmt.Println("Hello, I am John ")
// }
// func (b BotEnglish) SayDay() {
// 	fmt.Println("Hello, I am John ")
// }
// func (b BotEnglish) SayBye() {
// 	fmt.Println("Hello, I am John ")
// }
