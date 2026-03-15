package main

import "fmt"

type User struct {
	UserName string
	Password string
}

type Login struct {
	Users map[string]User
}

func main() {
	startLogin()
}

func startLogin() {

	login := Login{
		Users: make(map[string]User),
	}

	for {
		choice := menuLogin()

		switch choice {
		case 1:
			registerUser(&login)

		case 2:
			loginUser(&login)

		case 3:
			viewUsers(&login)

		case 4:
			fmt.Println("Exiting Program...")
			return

		default:
			fmt.Println("Invalid Choice")
		}
	}
}

func menuLogin() int {

	var choice int

	fmt.Println("\n===== LOGIN SYSTEM =====")
	fmt.Println("1. Register User")
	fmt.Println("2. Login User")
	fmt.Println("3. View Users")
	fmt.Println("4. Exit")
	fmt.Print("Enter Choice: ")

	fmt.Scanln(&choice)

	return choice
}

func registerUser(login *Login) {

	fmt.Println("\n===== USER REGISTRATION =====")

	var userName string
	var password string

	fmt.Print("Enter Username: ")
	fmt.Scanln(&userName)

	if userName == "" {
		fmt.Println("Username cannot be empty")
		return
	}

	_, exist := login.Users[userName]

	if exist {
		fmt.Println("User already exists")
		return
	}

	fmt.Print("Enter Password: ")
	fmt.Scanln(&password)

	if password == "" {
		fmt.Println("Password cannot be empty")
		return
	}

	user := User{
		UserName: userName,
		Password: password,
	}

	login.Users[userName] = user

	fmt.Println("User Registered Successfully")
}

func loginUser(login *Login) {

	fmt.Println("\n===== USER LOGIN =====")

	var userName string
	var password string

	fmt.Print("Enter Username: ")
	fmt.Scanln(&userName)

	value, exist := login.Users[userName]

	if !exist {
		fmt.Println("User does not exist")
		return
	}

	fmt.Print("Enter Password: ")
	fmt.Scanln(&password)

	if value.Password == password {
		fmt.Println("Login Successful")
	} else {
		fmt.Println("Invalid Password")
	}
}

func viewUsers(login *Login) {

	fmt.Println("\n===== REGISTERED USERS =====")

	if len(login.Users) == 0 {
		fmt.Println("No users registered")
		return
	}

	for key, value := range login.Users {
		fmt.Println("Username:", key, "| Password:", value.Password)
	}
}