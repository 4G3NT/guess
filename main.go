package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func genRandNum(num int) int {
	seedNum := time.Now().UnixNano()
	rand.Seed(seedNum)
	return rand.Intn(num) + 1
}

func input(prompt string, rdr *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := rdr.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	argument := input("How much range would you like to put?\n\nE.g. 100 will do 1-100: ", reader)
	arg, err := strconv.Atoi(argument)

	if err != nil {
		color.Red("Error, Please enter a number instead of letter(s)")
	}

	secret := genRandNum(arg)
	fmt.Printf("I've randomly generated a random number 1-%v.\n", arg)
	attempsStr := input("How many attempts would you like?: ", reader)

	attemptsConverted, err := strconv.Atoi(attempsStr)

	if err != nil {
		color.Red("Error, Please enter a number instead of letter(s)")
	}

	for tries := 0; tries < attemptsConverted; tries++ {
		fmt.Printf("You have %v guesses left.\n", attemptsConverted-tries)

		inp := input("Enter your guess: ", reader)

		guess, err := strconv.Atoi(inp)

		if err != nil {
			color.Red("Error, Please enter a number instead of letter(s)")
		}

		if guess < secret {
			color.Yellow("Oops. Your guess was too LOW.")
		} else if guess > secret {
			color.Yellow("Oops. Your guess was too HIGH.")
		} else {
			color.Green("\n\nYou guessed the number!, The number was %v.\n", secret)
			break
		}
	}
}
