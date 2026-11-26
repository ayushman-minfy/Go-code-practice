package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	words := strings.Fields(string(data))
	fmt.Println("Total no of Words in the file: ", len(words))
}
