package main

import (
	"fmt"
)

type AgeError struct {
	Age int
}

func (e AgeError) Error() string {
	return fmt.Sprintf("Age %d is not allowed. Must be 18 or older.", e.Age)
}

func checkAge(age int) error {

	if age < 18 {
		return AgeError{Age: age}
	}

	return nil
}

func main() {

	err := checkAge(16)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Age is valid")
}