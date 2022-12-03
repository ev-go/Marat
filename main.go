package main

import "fmt"

func main() {
	fmt.Println("Добро пожаловать это калькулятор:")
	fmt.Println("Выберете операцию где \n+ это сложение\n- это вычитание\n* это умножение\n/ это деление")
	var operation string
	fmt.Scan(&operation)
	var firstnumber, secondnumber int

	fmt.Println("Введите два числа:\nВведите первое чило:")
	fmt.Scan(&firstnumber)
	fmt.Println("Введите второе число:")
	fmt.Scan(&secondnumber)

	if operation == "+" {
		resultOfadding := adding(firstnumber, secondnumber)
		fmt.Println("Сложение  a и b:", resultOfadding)
	}
	if operation == "-" {
		resultOfminus := minus(firstnumber, secondnumber)
		fmt.Println("ВЫчитание чисел a и b:", resultOfminus)
	}
	if operation == "*" {
		resultOfmulti := multi(firstnumber, secondnumber)
		fmt.Println("Произведение числа a и b:", resultOfmulti)
	}

	if operation == "/" {
		resultOfdiv := div(firstnumber, secondnumber)
		fmt.Println("Деление чисел a и b:", resultOfdiv)
	}

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
