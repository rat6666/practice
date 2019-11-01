package bot

import (
	"fmt"
	"os"
	"time"
)

type BotEnglish struct {
	Name   string
	Appeal string
}

func (b BotEnglish) SayHello() {
	fmt.Printf("%v: Hello, I am %v\n", b.Name, b.Name)
}

func (b BotEnglish) SayTime() {
	location := "Europe/London"
	t := time.Now()
	loc, err := time.LoadLocation(location)
	if err != nil {
		fmt.Println(err)
	}
	t = t.In(loc)
	fmt.Printf("%v: Now in London is %v\n", b.Name, t.Format("15:04"))
}
func (b BotEnglish) SayDate() {
	t := time.Now()
	fmt.Printf("%v: Today is %v %v %v\n", b.Name, t.Month(), t.Day(), t.Year())
}
func (b BotEnglish) SayDay() {
	t := time.Now()
	fmt.Printf("%v: Today is %v\n", b.Name, t.Weekday())
}
func (b BotEnglish) SayBye() {
	fmt.Printf("%v: Bye\n", b.Name)
	os.Exit(1)
}
