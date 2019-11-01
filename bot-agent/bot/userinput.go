package bot

import (
	"fmt"
	"strings"
)

var language string

func UserInput(prefix string) string {
	for {
		fmt.Printf("%v:", prefix)
		_, err := fmt.Scanln(&language)
		if err != nil || language == "" {
			fmt.Print("Wrong input:", err, "\n")
		} else {
			return strings.ToLower(language)
		}
		fmt.Println("Try again/Попробуйте снова")
	}
}
