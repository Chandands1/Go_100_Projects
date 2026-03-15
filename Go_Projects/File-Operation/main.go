package main


import ("fmt"
         "os"
		"encoding/json")


type Student struct{
	Name string `json:name`
	Age int `json:age`
	Skill string `json:skill`
}	




func startJson(){

	student := Student{}
	student.Name = "Chadan"
	student.Age = 14
	student.Skill = "Java"

	jsonData, err := json.Marshal(student)

	if err != nil{
		fmt.Println("Error found")
		return
	}
	err = os.WriteFile("user.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("JSON written successfully to user.json")

    student2 := Student{}

	data,  err := os.ReadFile("user.json")
	fmt.Println(string(data))

	if err != nil{
		fmt.Println("Error")
		
		return
	}

	err =  json.Unmarshal(data, &student2)
	if err != nil{
		fmt.Println("Error", err)
		
		return
	}

	fmt.Println(student2.Name, student2.Age, student2.Skill)

}




func main(){

	startJson()

}		
