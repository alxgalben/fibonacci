package main

import "fmt"

func next_fibonacci(a, b *int) (response int) {
	response = *a + *b
	*a = *b
	*b = response

	return
}

var blacklist = []int{1, 5, 13, 21}

func checkListAndAppend(list []int, val int) []int {
	for _, value := range blacklist {
		if val == value {
			return list
		}
	}
	list = append(list, val)
	return list
}

func main() {

	/*var blacklistSize int
	fmt.Printf("Introduceti marimea blacklist-ului: ")
	fmt.Scanln(&blacklistSize)
	var blacklist = make([]int, blacklistSize)
	for i := 0; i < blacklistSize; i++ {
		fmt.Printf("Introduceti %d element: ", i)
		fmt.Println("")
		fmt.Scanln(&blacklist[i])
	}
	fmt.Print("Blacklist-ul final este: ", blacklist)
	*/

	var userIndex int
	var userRange int
	finalArray := make([]int, 0)
	element_1 := 0
	element_2 := 1

	fmt.Println("De la care index sa incepem? ")
	fmt.Scan(&userIndex)
	fmt.Println("Ce lungime sa fie? ")
	fmt.Scan(&userRange)
	fmt.Println("Sirul final este :")

	finalArray = checkListAndAppend(finalArray, 0)
	finalArray = checkListAndAppend(finalArray, 1)
	suma_totala := userIndex + userRange

	for len(finalArray) < suma_totala {
		finalArray = checkListAndAppend(finalArray, next_fibonacci(&element_1, &element_2))
	}

	fmt.Println(finalArray[userIndex:suma_totala])
}
