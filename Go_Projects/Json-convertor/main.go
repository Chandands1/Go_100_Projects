package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	City  string `json:"city"`
}

func main() {

	// JSON data
	jsonData := `{
		"name": "Chandan",
		"age": 25,
		"email": "chandan@gmail.com",
		"city": "Bangalore"
	}`

	// Create struct variable
	var user User

	// Convert JSON to Struct
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print struct values
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("Email:", user.Email)
	fmt.Println("City:", user.City)

}