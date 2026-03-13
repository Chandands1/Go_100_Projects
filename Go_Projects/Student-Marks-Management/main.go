package main

import (
	"fmt"
)

type Student struct {
	Name       string
	RollNumber string
	Total      float64
	Average    float64
	Grade      rune
}

func startStudent() {
	std1 := createStudent()
	enterMarks(&std1)
}

func createStudent() Student {
	var std Student

	fmt.Print("Enter the student name: ")
	fmt.Scanln(&std.Name)

	fmt.Print("Enter the Roll number: ")
	fmt.Scanln(&std.RollNumber)

	return std
}

func enterMarks(std *Student) {
	var math, science, social float64

	fmt.Println("Enter your marks for each subject")

	fmt.Print("Enter Math marks: ")
	fmt.Scanln(&math)

	fmt.Print("Enter Science marks: ")
	fmt.Scanln(&science)

	fmt.Print("Enter Social marks: ")
	fmt.Scanln(&social)

	std.Total = calculateTotal(math, science, social)

	std.Average = calculateAverage(std.Total)

	std.Grade = calculateGrade(std.Average)

	displayResult(std)
}

func calculateTotal(math, science, social float64) float64 {
	return math + science + social
}

func calculateAverage(total float64) float64 {
	return total / 3
}

func calculateGrade(average float64) rune {

	if average >= 90 {
		return 'A'
	} else if average >= 80 {
		return 'B'
	} else if average >= 70 {
		return 'C'
	} else if average >= 60 {
		return 'D'
	} else {
		return 'F'
	}

}

func displayResult(std *Student) {

	fmt.Println("\n----- STUDENT RESULT -----")
	fmt.Println("Name       :", std.Name)
	fmt.Println("RollNumber :", std.RollNumber)
	fmt.Println("Total Marks:", std.Total)
	fmt.Println("Average    :", std.Average)
	fmt.Println("Grade      :", string(std.Grade))

}

func main() {
	startStudent()
}