package main

import (
	"bufio"
	"fmt"
	"github.com/LoyalEnv0y/Mors-Encode-Decode/morse"
	"os"
)

func main() {
	var (
		userChose string
		userInput string
	)

	for {
		fmt.Println("******** WELCOME TO MORSE CODE TRANSLATOR ********")
		fmt.Printf("Type 1 for Latin to Morse\nType 2 for Morse to Latin\n")
		_, err := fmt.Scanf("%s\n", &userChose)
		if err != nil {
			fmt.Printf("%s\n Please type again", err)
			continue
		}
		break
	}

	fmt.Println("Please type the massage you want to translate")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput = scanner.Text()

	switch userChose {
	case "1":
		s, err := morse.Encode(userInput)
		if err != nil {
			fmt.Println("Error!", err)
		} else {
			fmt.Println(s)
		}

	case "2":
		s, err := morse.Decode(userInput)
		if err != nil {
			fmt.Println("Error!", err)
		} else {
			fmt.Println(s)
		}
	}

}
