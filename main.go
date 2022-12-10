package main

import "fmt"

func main() {
	//var n, q, w int
	//fmt.Scan(&n)
	//if n < 4 {
	//	panic(n)
	//}
	//a := make([]int, n)
	//for i := 0; i < n; i++ {
	//	fmt.Scan(&q)
	//	a[i] = q
	//}
	//w = a[3]
	//fmt.Println(w)
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	var result = array[0]
	for i := 0; i < 5; i++ {

		if result > array[i] {
			result = result
		} else {
			result = array[i]
		}

	}
	fmt.Println(result)
}

//	baseArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//	baseSlice := baseArray[5:8]
//	fmt.Printf("Слайс: %v\n", baseSlice)
//	baseSlice = append(baseSlice, 10)
//	fmt.Printf("Слайс после апенда: %v\n", baseSlice)
//	fmt.Printf("исходный аррэй массив: %v\n", baseArray)
//	pointer := fmt.Sprintf("%p", baseSlice)
//	fmt.Printf("адрес в поинтере: %v\n", pointer)
//	fmt.Printf("Массив: %v\n", baseArray)
//	fmt.Printf("Срез длиной %d и емкостью %d: %v\n", len(baseSlice), cap(baseSlice), baseSlice)
//	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice))
//
//	b := make([]int, 3, 3)
//	fmt.Println(b)
//
//}

//	fmt.Println("Добро пожаловать это калькулятор:")
//	fmt.Println("Выберете операцию где \n+ это сложение\n- это вычитание\n* это умножение\n/ это деление")
//	var operation string
//	fmt.Scan(&operation)
//	var firstnumber, secondnumber int
//
//	fmt.Println("Введите два числа:\nВведите первое чило:")
//	fmt.Scan(&firstnumber)
//	fmt.Println("Введите второе число:")
//	fmt.Scan(&secondnumber)
//
//	if operation == "+" {
//		resultOfadding := adding(firstnumber, secondnumber)
//		fmt.Println("Сложение  a и b:", resultOfadding)
//	}
//	if operation == "-" {
//		resultOfminus := minus(firstnumber, secondnumber)
//		fmt.Println("ВЫчитание чисел a и b:", resultOfminus)
//	}
//	if operation == "*" {
//		resultOfmulti := multi(firstnumber, secondnumber)
//		fmt.Println("Произведение числа a и b:", resultOfmulti)
//	}
//
//	if operation == "/" {
//		resultOfdiv := div(firstnumber, secondnumber)
//		fmt.Println("Деление чисел a и b:", resultOfdiv)
//	}
//
//}
//func adding(a int, b int) (z int) {
//	z = a + b
//	return z
//}
//
//func multi(a int, b int) (z int) {
//	z = a * b
//	return z
//}
//func div(a int, b int) (z int) {
//	z = a / b
//	return z
//}
//func minus(a int, b int) (z int) {
//	z = a - b
//	return z
//}
