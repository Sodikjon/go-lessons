package main

import "fmt"

func main() {
	// Напишите программу, которая запрашивает у пользователя ввод положительного числа и повторяет запрос,
	// пока не будет введено положительное число.

	var a int

	for {
		fmt.Println("Введите положительное число: ")
		fmt.Scanln(&a)
		if a > 0 {
			fmt.Println("Урааа ")
			break
		}
	}

}
