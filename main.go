package main

import "fmt"


func main(){

	var x, y float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&x)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&y)

	fmt.Print("Выберите оператора ( + - * / ) : ")
	fmt.Scanln(&operator)

	switch operator{

	case "+" :
		fmt.Printf("%.3f %s %.3f = %.3f", x, operator, y, x + y)
	case "-" :
		fmt.Printf("%.3f %s %.3f = %.3f", x, operator, y, x - y)	
	case "*" :
		fmt.Printf("%.3f %s %.3f = %.3f", x, operator, y, x * y)
	case "/" :
		if(y == 0){
		fmt.Println("Ошибка! Обнаружено деление на 0.")
		break;
		}
		fmt.Printf("%.3f %s %.3f = %.3f", x, operator, y, x / y)
	default:
		fmt.Println("Ошибка! Введен неверный оператор.")
	}
}