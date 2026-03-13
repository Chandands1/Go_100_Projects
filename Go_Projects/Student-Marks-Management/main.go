package main

import "fmt"

type Marks struct {
	Subject string
	Mark    float64
}

type Student struct {
	name    string
	s_id    int
	marks   []Marks
	total   float64
	average float64
	grade   string
}

func CreateStudent() {
	s1 := Student{}

	fmt.Println("Enter Student Name:")
	fmt.Scanln(&s1.name)

	fmt.Println("Enter Student ID:")
	fmt.Scanln(&s1.s_id)

	fmt.Println("Enter number of subjects:")
	var num int
	fmt.Scanln(&num)

	CreateMarks(&s1, num)
	TotalMarks(&s1)
	AverageMarks(&s1)
	Grade(&s1)
	Display(&s1)
}

func CreateMarks(s1 *Student, num int) {

	for i := 0; i < num; i++ {

		var m Marks

		fmt.Println("Enter Subject Name:")
		fmt.Scanln(&m.Subject)

		fmt.Println("Enter Marks:")
		fmt.Scanln(&m.Mark)

		if m.Mark > 100 || m.Mark < 0 {
			fmt.Println("Invalid marks entered. Try again.")
			i--
			continue
		}

		s1.marks = append(s1.marks, m)
	}
}

func TotalMarks(s1 *Student) {

	total := 0.0

	for _, m := range s1.marks {
		total += m.Mark
	}

	s1.total = total
}

func AverageMarks(s1 *Student) {

	s1.average = s1.total / float64(len(s1.marks))
}

func Grade(s1 *Student) {

	switch {

	case s1.average >= 90:
		s1.grade = "A"

	case s1.average >= 80:
		s1.grade = "B"

	case s1.average >= 70:
		s1.grade = "C"

	case s1.average >= 60:
		s1.grade = "D"

	case s1.average >= 50:
		s1.grade = "E"

	default:
		s1.grade = "F"
	}
}

func Display(s1 *Student) {

	fmt.Println("\n------ Student Details ------")

	fmt.Println("Student Name:", s1.name)
	fmt.Println("Student ID:", s1.s_id)

	fmt.Println("\nSubjects and Marks:")

	for _, m := range s1.marks {
		fmt.Println("Subject:", m.Subject, "Marks:", m.Mark)
	}

	fmt.Println("\nTotal Marks:", s1.total)
	fmt.Println("Average:", s1.average)
	fmt.Println("Grade:", s1.grade)
}

func StartStudent() {

	fmt.Println("Welcome to Student Information System")
	CreateStudent()
}

func main() {
	StartStudent()
}