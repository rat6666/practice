package bot

import (
	"fmt"
	"os"
)

func Start() {
	fmt.Println("Input English for english bot or Введите Russian для русского бота.")
	var language string
	_, err := fmt.Scanln(&language)
	if err != nil {
		fmt.Print("Something wrong:", err)
		os.Exit(1)
	}
	for language != "" {

	}
	fmt.Println(language)

}
