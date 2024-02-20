package main

import "fmt" //Import to be able to print the result to the console

func calculate(num1 int, operator rune, num2 int) int {
	switch operator {
	case '+':
		return num1 + num2
	case '-':
		return num1 - num2
	default:
		return 0
	}
}

func main() {
	num1 := 5
	num2 := 7
	operator := '-'

	result := calculate(num1, operator, num2)
	fmt.Println(result)
}
