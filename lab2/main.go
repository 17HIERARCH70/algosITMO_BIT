package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/17HIERARCH70/betonicSort"
)

func main() {
	// CLI UI
	fmt.Println("Betonic sort")
	fmt.Println("-------------")
	fmt.Println("Choose input source:")
	fmt.Println("1. Standard input (keyboard)")
	fmt.Println("2. File")

	reader := bufio.NewReader(os.Stdin)

	var choice int
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		log.Fatal("Error reading choice: ", err)
	}

	var input string

	switch choice {
	case 1: // read from keyboard
		fmt.Println("Enter array of length 2^n:")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading from keyboard: ", err)
		}
		input = strings.TrimSpace(input)
	case 2: // read from file
		fmt.Print("Enter file path (data should be array of length 2^n): ")
		filePath, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal("Error reading file path: ", err)
		}
		filePath = strings.TrimSpace(filePath) // remove trailing newline character

		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal("Error opening file: ", err)
		}
		// close file on exit
		defer func() {
			if err := file.Close(); err != nil {
				log.Fatal("Error closing file: ", err)
			}
		}()
		inputBytes, err := io.ReadAll(file)
		if err != nil {
			log.Fatal("Error reading file: ", err)
		}
		input = string(inputBytes)

	default: // invalid choice
		log.Fatal("Invalid choice")
	}

	if len(input) == 0 {
		fmt.Println("Empty array. Exit.")
		return
	}

	// sort array
	sortedArr, err := betonicSort.BetonicSort(input)
	if err != nil {
		log.Fatal("Error sorting array: ", err)
	}

	// print sorted array
	fmt.Println("Sorted array:", sortedArr)
}
