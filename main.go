package main

import "fmt"

func main() {
	add()
}

func add() {
	var a, b int
	fmt.Println("Введите два числа:\nВведите первое чило:")
	fmt.Scan(&a)
	fmt.Println("Введите второе число:")
	fmt.Scan(&b)
	fmt.Println("Произведение числа a и b:", a*b)
}
