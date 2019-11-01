package bot

import (
	"fmt"
	"os"
	"time"
)

type BotRussian struct {
	Name   string
	Appeal string
}

var day = map[int]string{
	1: "Понедельник",
	2: "Вторник",
	3: "Среда",
	4: "Четверг",
	5: "Пятница",
	6: "Суббота",
	7: "Воскресенье",
}

var month = map[int]string{
	1:  "Января",
	2:  "Февраля",
	3:  "Марта",
	4:  "Апреля",
	5:  "Мая",
	6:  "Июня",
	7:  "Июля",
	8:  "Августа",
	9:  "Сентября",
	10: "Октября",
	11: "Ноября",
	12: "Декабря",
}

func (b BotRussian) SayHello() {
	fmt.Printf("%v: Привет, я %v\n", b.Name, b.Name)
}
func (b BotRussian) SayTime() {
	name := "Europe/Minsk"
	t := time.Now()
	loc, err := time.LoadLocation(name)
	if err != nil {
		fmt.Println(err)
	}
	t = t.In(loc)
	fmt.Printf("%v: Cейчас в Беларуси %v\n", b.Name, t.Format("15:04"))
}
func (b BotRussian) SayDate() {
	t := time.Now()
	fmt.Printf("%v: Сегодня %v %v %v года\n", b.Name, t.Day(), month[int(t.Month())], t.Year())
}
func (b BotRussian) SayDay() {
	t := time.Now()
	fmt.Printf("%v: Сегодня %v\n", b.Name, day[int(t.Weekday())])
}
func (b BotRussian) SayBye() {
	fmt.Printf("%v: Пока\n", b.Name)
	os.Exit(1)
}
