package main

import (
	"fmt"
	//"runtime/trace"
	"slices"
	"unicode"

	"strings"
)

func ReverseAnArray(arr [7]int)[7]int{

	reverseArray := [7]int{}

	for i,_ := range arr{
		reverseArray[6-i] = arr[i]
	}

	return reverseArray
	


}


func LinearSearch(arr [7]int, element int){
	var flag bool = false

	for i,_:= range arr{
		if arr[i] == element{
			flag = true
		}
	}
	if flag{
		fmt.Println("The element is present", element)
	}else{
		fmt.Println("Element is not present")
	}

}

func BubbleSort(arr [7]int){
	for i,_:= range arr{
		for j,_ := range arr{
			if	arr[i]<arr[j]{
		
		arr[i], arr[j] = arr[j], arr[i]
	}
	}
		}
	
	fmt.Println(arr)
}

func SelectionSort(arr [7]int){
	

	for i := 0; i < len(arr)-1; i++ {

		minIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	fmt.Println(arr)
}

func InsertionSorting(arr [7]int){
	for i := 1; i < len(arr); i++ {

		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}

	fmt.Println(arr)

}
func MergeTwoArrays(arr [7]int, arr2 [5]int){
	merged := append(arr[:],arr2[:]...)

	fmt.Println(merged)
	
}

func SecondLargest(arr []int){
	slices.Sort(arr)
	fmt.Println(arr)

	length := len(arr)

	if arr[length-1] > arr[length-2]{
		fmt.Println(arr[length-2])
	}else{
		fmt.Println(arr[length-3])
	}

}

func RemoveDuplicate(arr []int){
	for i:=0;i< len(arr)-1; i++{
		for j:=i+1; j<len(arr); j++{
		if	arr[i] ==arr[j]{
			slices.Delete(arr,j,j+1)
		}

		}
		


	}
	fmt.Println(arr)
}

func Frequency(arr []int){
    freq := make(map[int]int)
	for _, v:= range arr{
		freq[v]++
	}
	fmt.Println(freq)
}

func MoveArray(arr []int,k int){
	n:= len(arr)

	for i:=0;i<k;i++{
		temp := arr[0]
		for j:= 0;j<len(arr)-1;j++{
			arr[j]= arr[j+1]
		}
	    arr[n-1]= temp
	}

	fmt.Println(arr)

	for i:=0;i<k;i++{
		temp := arr[n-1]
		for j :=n-1;j >0;j--{
			arr[j]=arr[j-1] 
		}
		arr[0] = temp
	}
	fmt.Println(arr)
}

func SplitArrayIntoEvenOrOdd(arr []int){

	var evenArray []int
	var oddArray []int

	for _, v := range arr{
		if v % 2 ==0{
			evenArray = append(evenArray, v)
		}else{
			oddArray = append(oddArray, v)
		}
	}
	fmt.Println(evenArray)
	fmt.Println(oddArray)

}

func MissingNumber(arr []int){ //1,2,3,4,6,7,8
	
	sum := 0
	length := len(arr)
	firstElement:= arr[0]
	lastElement := arr[length-1]
	OriginalSum :=0
	for i:=firstElement; i <= lastElement; i++{
		OriginalSum += i

	}
     	fmt.Println(OriginalSum)

	for i:= 0;i<lastElement-1;i++{
		sum = sum + arr[i]
		fmt.Println(i)
	}
	fmt.Println(sum)

	fmt.Println(OriginalSum-sum)
	

}

func CountVowelsAndConsonant(str string){
	vowels :=0
	consonants :=0
	space :=0
	for _,v:= range str{
		if string(v) == "a"|| string(v)=="e"||string(v)=="i"||string(v)=="o"||string(v)=="u"{
           vowels++
		}else if string(v)==" "{
          space++
		}else{
			consonants++
		}
	}
	fmt.Println("Vowels:", vowels)
	fmt.Println(space)
	fmt.Println("Consonants:", consonants)
}

func ReversString(str string){
	originalString := str
	reverseString := ""
	for i,_:=  range str{
	 	reverseString += string(str[len(str)-1-i])
	}

	fmt.Println(reverseString)
	originalString = strings.ToLower(originalString)
	reverseString = strings.ToLower(reverseString)
	if originalString == reverseString{
		fmt.Println("String is Pallindrome")
	}else{
		fmt.Println("Not Palindrom")
	}
}


func CountWords(str string){
     count := 0
	for _, v := range str{
	if	unicode.IsLetter(v){
		count++
	}
	}
	fmt.Println(count)
}


func FrequencyCharacter(str string){

	freq := make(map[string]int)
	for _,v := range str{
		
		a :=string(v)
		if a == " "{
			continue
		}
		freq[a]++
	
	}
	fmt.Println(freq)
}

func RemoveDuplicateChar(str string){
	
    seen := make(map[rune]bool)
	result := ""

	for _, ch := range str{
		if !seen[ch]{
			seen[ch] = true
			result = result+ string(ch)
		}
	}
	fmt.Println(result)
}

func FindFirstNonRepeatingChar(str string){
	 count := make(map[rune]int)
	 
	 for _, ch := range str{
		count[ch]++
	 }

	 for _, ch := range str{
		if count[ch] == 1{
			fmt.Println("First non-repeating character:", string(ch))
			return
		}
	 }
	 fmt.Println("No non-repeating character found")

}


func Anagram(str1, str2 string)string{
	var uni1 rune
	var uni2 rune
	if len(str1) != len(str2){
		return "not equal"
	}
	for _, ch := range str1{
		uni1 = uni1 + ch
	}
	for _, ch := range str2{
		uni2 = uni2 + ch
	}
	if uni1 == uni2{
		return "Anagram"
	}
	return "Not a Anagram"


	
	
}

func CapitalizeFirstLetter(str string){
  words := strings.Fields(str)
  for i, word := range words{
	if len(word)>0{
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		words[i] = string(runes)
	}
  }
  strings.Join(words, " ")
  fmt.Println(words)
   	
}

func LongestWordInAString(str string){
	maxLenght := ""
	 words := strings.Fields(str)
	 for _, word := range words{
		if len(word)> len(maxLenght){
			 maxLenght = word
		}
		
		



	 }
	 fmt.Println(maxLenght)



}





func main(){
	//arr := []int{1,2,3,4,4,6,7,8}
	
	//LinearSearch(arr,60)

	//BubbleSort(arr)

	//SelectionSort(arr)

	//InsertionSorting(arr)
	//MergeTwoArrays(arr ,arr2)

	//SecondLargest(arr)
	//RemoveDuplicate(arr)

	//Frequency(arr)

	//MoveArray(arr,3)

	//SplitArrayIntoEvenOrOdd(arr)

	//MissingNumber(arr)
	// var str1 string = "Hello"

	// var str2 string = "Hello"

	//CountVowelsAndConsonant(str)

	//ReversString(str)

	//CountWords(str)

	//FrequencyCharacter(str)

	//RemoveDuplicateChar(str)

//	FindFirstNonRepeatingChar(str)

//fmt.Println(Anagram(str1,str2))
var str string = "hello thisthat is me"

//CapitalizeFirstLetter(str)
LongestWordInAString(str)










   
	
}
