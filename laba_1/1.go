package main

import (
	"errors"
	"fmt"
)



func hello(name string) string {
	return "Привет, " + name + "!"
}



func printEven(a, b int64) error {
	if a > b {
		return errors.New("левая граница больше правой")
	}
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	return nil
}


	
func apply(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("действие не поддерживается")
	}
}

func main() {
	fmt.Println(hello("name"))

	err := printEven(2, 15)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	err = printEven(15, 2)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	result, err := apply(2, 7, "+")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %f\n", result)
	}

	result, err = apply(5, 10, "*")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %f\n", result)
	}

	result, err = apply(3, 5, "#")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %f\n", result)
	}
}
