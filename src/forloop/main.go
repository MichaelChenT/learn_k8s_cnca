package main

import "fmt"

func main() {
	newString := [5]string{"I", "am", "stupid", "and", "weak"}
	for index, _ := range newString {
		newString[2] = "smart"
		newString[4] = "strong"
		fmt.Printf("%s\t", newString[index])
	}
}
