package main

import (
	"fmt"
	"labPP3/mathutils"
	"labPP3/stringutils"
)

func main() {
	// 1-2 Ввод числа и вычисление его факториала
	var num int
	fmt.Print("Введите число для вычисления факториала: ")
	fmt.Scan(&num)

	factorial := mathutils.Factorial(num)
	fmt.Printf("Факториал числа %d равен %d\n", num, factorial)

	// 2 Переворот строки
	var str string
	fmt.Print("Введите строку для переворота: ")
	fmt.Scan(&str)

	stringutils.ReverseString(str)

	// 3 Создание и вывод массива
	array := [5]int{}
	fmt.Println("Введите 5 целых чисел:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&array[i])
	}

	fmt.Println("Введенные числа:")
	for _, value := range array {
		fmt.Println(value)
	}

	// 4 Создание среза и работа с ним
	slice := array[:]
	fmt.Println("Созданный срез:", slice)

	// Добавление элементов (требуется использовать `append`)
	slice = append(slice, 100)
	fmt.Println("Срез после добавления элемента:", slice)

	// Удаление элемента
	slice = append(slice[:2], slice[3:]...)
	fmt.Println("Срез после удаления элемента:", slice)

	// 5 Работа со срезом строк
	strSlice := []string{"apple", "banana", "pineapple", "cherry"}
	maxLenStr := strSlice[0]
	for _, s := range strSlice {
		if len(s) > len(maxLenStr) {
			maxLenStr = s
		}
	}
	fmt.Printf("Самая длинная строка: %s\n", maxLenStr)
}
