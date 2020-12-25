package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	//List of words for the game to use
	var wordlist = [10]string{
		"Kangaroo",
		"Elephant",
		"Zoo",
		"Monkey",
		"Giraffe",
		"Lion",
		"Bear",
		"Parrot",
		"Cheeta",
		"Jaguar",
	}

	//Allows different randome item to get selected every time
	rand.Seed(time.Now().Unix())

	var play string = "y"
	var randomIndex int
	var chosen_word string
	var chances int = 5
	var letter string
	var chars []string
	var lenWord int
	for play == "y" {
		//Chooses word for instance of game
		randomIndex = rand.Intn(len(wordlist))
		chosen_word = strings.ToLower(wordlist[randomIndex])

		//Character spliter
		chars = strings.Split(chosen_word, "")
		var length int = len(chars)
		var guess []string
		for i := 0; i < length; i++ {
			guess = append(guess, "_")
		}
		lenWord = 0
		chances = 5
		for chances > 0 {
			//User Interface
			//fmt.Println(chars, len(chars))
			fmt.Printf("You have %v chances left\n", chances)
			fmt.Printf("%v\n", guess)
			fmt.Scanln(&letter)
			letter = strings.ToLower(letter)

			//Game Engine

			//Character checker
			var missed bool = true
			for i := 0; i < length; i++ {
				if letter == chars[i] {
					guess[i] = chars[i]
					lenWord++
					missed = false
				}

			}
			if lenWord == len(chosen_word) {
				fmt.Println("Congrats, You Won!")
				break
			}
			if missed {
				chances--
			}

		}
		fmt.Println("Game Over: Would you like to play again? (y/n)")
		fmt.Scanln(&play)
		play = strings.ToLower(play)
	}
}
