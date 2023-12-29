package main

import (
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func readTextFile() ([]string, error) {

	//Pass the file name to the ReadFile() function from
	//the util package to get the content of the file.

	content, err := os.ReadFile("russian_nouns.txt")

	// defer .Close()
	// Check whether the 'error' is nil or not. If it
	//is not nil, then print the error and exit.
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// convert the content into a string
	str := string(content)

	Arr := strings.Split(str, "\n")
	//Print the string str
	return Arr, nil
}

func findMatch(arr []string, letter string) []string {

	newArr := []string{}

	for _, val := range Arr {
		if strings.Contains(val, letter) {
			newArr = append(newArr, val)
		}
	}

	return newArr
}

func shrinkByLen(arr []string, a int) []string {

	newArr := []string{}

	for _, val := range Arr {
		numVal := utf8.RuneCountInString(val)
		if numVal <= a {
			newArr = append(newArr, val)
		}
	}

	return newArr

}

func Reset() []string {
	Arr, _ := readTextFile()
	return Arr
}

// func letterOrNumber(s string)  {

// }

// func main() {
// 	arr := readTextFile()

// 	fmt.Println(len(arr))

// 	endProgram := false

// 	for !endProgram {
// 		var letter string
// 		fmt.Println("Если хотите закончить программу, напишите - конец")
// 		fmt.Println("Если вы хотите ограничить длину поиска, напишите - число")
// 		fmt.Print("Enter letter: ")
// 		fmt.Scanln(&letter)
// 		if letter == "конец" {
// 			endProgram = true
// 			break
// 		}
// 		if letter == "число" {
// 			var a int
// 			fmt.Print("Введите число: ")
// 			fmt.Scanln(&a)
// 			arr = shrinkByLen(arr, a)
// 			fmt.Println(arr)
// 			fmt.Println(len(arr))
// 		}

// 		if letter != "число" && letter != "конец" {
// 			arr = findMatch(arr, letter)
// 			fmt.Println(arr)
// 			fmt.Println(len(arr))
// 		}

// 		// Problem that I run in to, is that the program will call
// 		// findMatch even then число is entered
// 		// that brakes the program...
// 		// arr = shrinkByLen(arr, len(letter))

// 	}

// }
