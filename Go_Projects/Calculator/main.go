package main

import "fmt"

func main() {
	startCalculator()
}

func startCalculator() {
	for {
		num1, num2 := getInput()
		choice := selectOperation()

		result := calculate(choice, num1, num2)
		fmt.Println("Result:", result)

		if !askRepeat() {
			fmt.Println("Thank you for using the calculator!")
			break
		}
	}
}

func getInput() (float64, float64) {
	var num1, num2 float64

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	return num1, num2
}

func selectOperation() string {
	var choice string

	fmt.Println("\nSelect Operation")
	fmt.Println("+  -> Addition")
	fmt.Println("-  -> Subtraction")
	fmt.Println("*  -> Multiplication")
	fmt.Println("/  -> Division")
	fmt.Println("%  -> Modulus")

	fmt.Print("Enter choice: ")
	fmt.Scanln(&choice)

	return choice
}

func calculate(choice string, a, b float64) float64 {
	switch choice {

	case "+":
		return addition(a, b)

	case "-":
		return subtraction(a, b)

	case "*":
		return multiplication(a, b)

	case "/":
		return division(a, b)

	case "%":
		return float64(modulus(int(a), int(b)))

	default:
		fmt.Println("Invalid operation")
		return 0
	}
}

func askRepeat() bool {
	var again string
	fmt.Print("\nDo you want to continue? (y/n): ")
	fmt.Scanln(&again)

	return again == "y" || again == "Y"
}

func addition(a, b float64) float64 {
	return a + b
}

func subtraction(a, b float64) float64 {
	return a - b
}

func multiplication(a, b float64) float64 {
	return a * b
}

func division(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Cannot divide by zero")
		return 0
	}
	return a / b
}

func modulus(a, b int) int {
	if b == 0 {
		fmt.Println("Error: Cannot perform modulus by zero")
		return 0
	}
	return a % b
}