package main

import (
	"fmt"

)

func sum(a int, b int)int{
	return a+b
}

func EvenChecker(number int)string{
	if number % 2==0{
		return "Even"
	}
	return "odd"
}

func SignChecker(num int){
	if num < 0{
		fmt.Println("Negative")
	}else if num > 0{
		fmt.Println("Positive")
	}else{
		fmt.Println("Zero")
	}
}


func MaxOf2Numbers(a int, b int)int{
	if a > b{
		return a
	}
	return b
}
func MaxOfNumbers(num ...int)int{
	max :=0

	for i, _:= range num{
		if num[i] > max{
			max = num[i]
			
		}
	}
	return max
}


func Swap2Numbers(a int, b int)(int, int){
	temp := a
	a = b
	b = temp

	return a, b
}

func Swap(a int, b int)(int, int){
     a,b = b,a

	return a , b
}

func LeapYear(year int)string{
	if year % 4 ==0 && year & 100 !=0 || year % 400==0{
		return "Leap year"
	}
	return "Not a leap year"
}

func PrintNumber(n int){
	for i:=1 ; i<=n ; i++{
		fmt.Println(i)
	}
}

func SumOFnNUmbers(num int)int{
	total:=0
	for  value := range num+1{
		total = total + value
		fmt.Println(total)
	}
	return total
}

func MultiplicationTable(number int){
	for i := 1; i<=10; i++{
		fmt.Println(number , "x", i, "=", number*i)
	}
}

func CountDigits(num int)int{
	count := 0
	for num !=0{
		
		num = num / 10
		count = count+1

	}
	return count
	
}

func ReverseNumber(number int)int{
	originalNumber := number
	var reverseNumber int

	for originalNumber !=0{
		digit := originalNumber % 10

		reverseNumber =reverseNumber* 10 + digit
		originalNumber = originalNumber/10
	}
	return reverseNumber

}

func Palindrome(number int){
	originalNumber := number
	var reverseNumber int

	for originalNumber !=0{
		digit := originalNumber % 10

		reverseNumber =reverseNumber* 10 + digit
		originalNumber = originalNumber/10
	}
	if(number == reverseNumber){
		fmt.Println("Palindrome")
	}else{
		fmt.Println("Not Palindrome")
	}

}

func Factorial(num int)int{
	
	if num ==0{
		return 1
	}


	return num * Factorial(num -1)
}

func FibonacciSeries(num int){
	a := 0
	b :=1
	
		fmt.Print(a , b, " ")

	for i:=3; i<= num; i++{
     sum := a+b
	fmt.Print(sum, " ")
	a = b
	b = sum
	}
	
}

func CheckPrimeNumber(number int){
	
	count := 0
	for i :=1 ;i <= number; i++{
	if	number % i == 0{
        
		count++
		
		
	}
		
	}
	if count == 2{
		fmt.Println("Prime")
	}else{
		fmt.Println("Not a Prime")
	}
}

func PrintPrime(num int) {

	for i := 2; i <= num; i++ {
		isPrime := true

		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Println(i)
		}
	}
}

func HCF(a int, b int){

	var num1 []int

	var num2 []int
	var hcf int



	for i := 1; i <= a; i++{
		if a % i ==0{
			num1 = append(num1, i)

		}
		


	}
	
	for i := 1; i <= b; i++{
		if b % i ==0{
			num2 = append(num2, i)

		}
		


	}
	fmt.Println("Factors of a:", num1)
	fmt.Println("Factors of b:", num2)
for _, v1 := range num1 {
		for _, v2 := range num2 {
			if v1 == v2 {
				hcf = v1
			}
		}
	}

	fmt.Println("HCF:", hcf)

	
}







func main(){

	//fmt.Println("Hello, World")
	// var number int
	// fmt.Println("Enter the number:")

	// fmt.Scanln(&number)

	//fmt.Println(sum(5,7))
	//fmt.Println(EvenChecker(71))

	//SignChecker(98)

	//fmt.Println("Maximum number is : ",MaxOf2Numbers(454,9))
	//fmt.Println("Maximum number is : ",MaxOfNumbers(454,9,573,345))

	//fmt.Println(Swap2Numbers(6,12))
	//fmt.Println(Swap(6,12))
	//fmt.Println(LeapYear(2028))
	//PrintNumber(545)

//	fmt.Println(SumOFnNUmbers(5))
//MultiplicationTable(6)



 //fmt.Println(CountDigits(1324))


 //fmt.Println(ReverseNumber(1234))

 //Palindrome(121)

 //fmt.Println(Factorial(5))

 //FibonacciSeries(10)

 //CheckPrimeNumber(9)

 //PrintPrime(3)

 HCF(12,18)












}