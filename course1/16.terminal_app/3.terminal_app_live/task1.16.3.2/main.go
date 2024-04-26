package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

// Функция для отображения меню
func displayMenu(menu string) {
	fmt.Println("\033[H\033[2J")
	fmt.Println(menu)
}

// Функция для обработки выбора пользователя
func handleChoice(char rune, key keyboard.Key) string {
	switch char {
	case '1':
		return submenu1()
	case '2':
		return submenu2()
	case 'q':
		os.Exit(0)
	case '\x00': // BACKSPACE
		if key == 127 {
			return mainMenu()
		}
	default:
		return mainMenu()
	}
	return ""
}

// Главное меню
func mainMenu() string {
	return "My menu:\n1. Submenu 1\n2. Submenu 2\nPress q to quit"
}

// Подменю 1
func submenu1() string {
	return "Submenu 1\nContent submenu 1\n"
}

// Подменю 2
func submenu2() string {
	return "Submenu 2\nContent submenu 2\n"
}

func main() {
	menu := mainMenu()
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		displayMenu(menu)

		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println(err)
			return
		}
		menu = handleChoice(char, key)
	}
}
