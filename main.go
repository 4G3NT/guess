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

func randNum(rng int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(rng) + 1
}

func input(prmpt string, rdr *bufio.Reader) string {
	fmt.Print(prmpt)
	input, _ := rdr.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	red := color.New(color.FgRed).FprintfFunc()

	reader := bufio.NewReader(os.Stdin)
askForRange:
	rngInput := input("How much range would you like to put?\n\ne.g. 100 will do 1 -> 100: ", reader)

	arg, err := strconv.Atoi(rngInput)

	if err != nil {
		red(os.Stderr, "error: \"%v\" is not a number.\n", rngInput)
		goto askForRange
	}

	rand := randNum(arg)
	fmt.Printf("I've randomly generated a random number 1 -> %v.\n", arg)
askForAttempts:
	attempts := input("How many attempts would you like?: ", reader)

	attemptsConverted, err := strconv.Atoi(attempts)

	if err != nil {
		red(os.Stderr, "error: \"%v\" is not a number.\n", attempts)
		goto askForAttempts
	}

	for tries := 0; tries < attemptsConverted; tries++ {
		fmt.Printf("You have %v guesses left.\n", attemptsConverted-tries)

		gss := input("Enter your guess: ", reader)

		guess, err := strconv.Atoi(gss)

		if err != nil {
			tries--
			color.Red("Error, Please enter a number instead of letter(s)")
			continue
		}

		if guess < rand {
			color.Yellow("Oops. Your guess was too LOW.")
		} else if guess > rand {
			color.Yellow("Oops. Your guess was too HIGH.")
		} else {
			color.Green("\n\nYou guessed the number!, The number was %v.\n", rand)
			os.Exit(0)
		}
	}
	fmt.Printf("Sadly you did't guess the number :(, The number was %v\n", rand)
}
