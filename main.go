package main

import "fmt"
import "strings"

func main() {
	var word string
	fmt.Scan(&word)
	word = strings.ToUpper(word)
	fmt.Print(word)
}
