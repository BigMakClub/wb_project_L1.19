package main

import (
	"bufio"
	"fmt"
	"os"


func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	letters := []rune(input)

	for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
		letters[i], letters[j] = letters[j], letters[i]
	}

	fmt.Println(string(letters))
}
