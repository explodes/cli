package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printMenu(menu ...string) {
	for index, item := range menu {
		fmt.Printf("%d) %s\n", index+1, item)
	}
}

func PresentMenu(ps1 string, menu ...string) int {
	printMenu(menu...)
	var selection int
	for {
		selection = PromptInt(ps1)
		if selection < 1 || selection > len(menu) {
			fmt.Println("invalid selection")
			continue
		}
		return selection - 1
	}
}

func PromptInt(ps1 string) int {
	for {
		input := PromptNonEmptyString(ps1)
		inputValue, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		return inputValue
	}
}

func PromptString(ps1 string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(ps1)
		input, _ := reader.ReadString('\n')
		return strings.TrimSuffix(input, "\n")
	}
}

func PromptNonEmptyString(ps1 string) string {
	for {
		input := PromptString(ps1)
		if input == "" {
			fmt.Println("invalid input")
			continue
		}
		return input
	}
}
