package mai

import ("fmt")

type Student struct{
	Name string
	RollNumber string
	Marks float64
	Total float64
	Average float64
	Grade rune
}

func startStudent(){
	std1 := createStudent()


}
func createStudent(std *Student){

	fmt.Print("Enter the student name:")
	fmt.Scanln(&std.Name)
	fmt.Print("Enter the Roll number:")
	fmt.Scanln(&std.RollNumber)
	fmt.Print("Enter the Total:")
	fmt.Scanln(&std.Total)


}


func main(){
	startStudent()

}