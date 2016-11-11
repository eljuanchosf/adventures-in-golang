/* Write a guessing game where the user has to guess a secret number.
After every guess the program tells the user whether their number was too large or too small.
At the end the number of tries needed should be printed.
I counts only as one try if they input the same number multiple times consecutively. */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type limits struct {
	bottom int
	top    int
}

func main() {
	limit := limits{1, 10}
	var guessStack []int
	numberToGuess := getRandomNumber(limit.bottom, limit.top)
	fmt.Printf("Hello! I have generated a randon number between %d and %d. Can you guess it?\n", limit.bottom, limit.top)
	fmt.Println("Press CTRL-C to exit at any time!")
	tries := 1
	for {
		userGuess := askForNumber("Please enter a number: ")
		if userGuess != numberToGuess {
			if !guessInStack(userGuess, guessStack) {
				guessStack = append(guessStack, userGuess)
				tries++
			}
		} else {
			fmt.Printf("Yesss!! You guessed in %d tries!!\n", tries)
			break
		}
	}
}

func getRandomNumber(bottomLimit int, topLimit int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(topLimit-bottomLimit) + bottomLimit
}

func guessInStack(guess int, stack []int) bool {
	for i := 0; i < len(stack); i++ {
		if stack[i] == guess {
			return true
		}
	}
	return false
}

func askForNumber(message string) int {
	var num int
	fmt.Printf(message)
	fmt.Scanf("%d", &num)
	return num
}
