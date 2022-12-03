package main

import "fmt"

func main() {
	var firstnumber, secondnumber int
	fmt.Println("Введите два числа:\nВведите первое чило:")
	fmt.Scan(&firstnumber)
	fmt.Println("Введите второе число:")
	fmt.Scan(&secondnumber)

	resultOfadding := adding(firstnumber, secondnumber)
	resultOfmulti := multi(firstnumber, secondnumber)
	resultOfdiv := div(firstnumber, secondnumber)
	resultOfminus := minus(firstnumber, secondnumber)
	fmt.Println("Произведение числа a и b:", resultOfmulti)
	fmt.Println("Сложение  a и b:", resultOfadding)
	fmt.Println("Деление чисел a и b:", resultOfdiv)
	fmt.Println("ВЫчитание чисел a и b:", resultOfminus)
}
func adding(a int, b int) (z int) {
	z = a + b
	return z
}

func multi(a int, b int) (z int) {
	z = a * b
	return z
}
func div(a int, b int) (z int) {
	z = a / b
	return z
}
func minus(a int, b int) (z int) {
	z = a - b
	return z
}
