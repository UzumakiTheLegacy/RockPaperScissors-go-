package main

import (
	"fmt"
	"math/rand"
)

var player string
var computer string
var result string

func generateComputerChoice() string {
	choices := []string{"rock", "paper", "scissors"}
	return choices[rand.Intn(3)]
}

func determineWinner(player string, computer string) string {
	if player == computer {
		return "tie"
	} else if player == "rock" && computer == "scissors" || player == "Rock" && computer == "Scissors" {
		return "player"
	} else if player == "scissors" && computer == "paper" || player == "Scissors" && computer == "Paper" {
		return "player"
	} else if player == "paper" && computer == "rock" || player == "Paper" && computer == "Rock" {
		return "player"
	} else {
		return "computer"
	}
}

func printWinner(winner string) {
	if winner == "player" {
		fmt.Println("You win!")
	} else if winner == "computer" {
		fmt.Println("You lose!")
	} else {
		fmt.Println("It's a tie!")
	}
}

func printPlayerChoice(player string) {
	fmt.Println("You chose " + player)
}

func printComputerChoice(computer string) {
	fmt.Println("The computer chose " + computer)
}

func printResult(result string) {
	if result == "player" {
		fmt.Println("You win!")
	} else if result == "computer" {
		fmt.Println("You lose!")
	} else {
		fmt.Println("It's a tie!")
	}
}

func printGame(player string, computer string, result string) {
	printPlayerChoice(player)
	printComputerChoice(computer)
	printResult(result)
}
func askUserForChoice() {
	fmt.Println("Choose rock, paper, or scissors:")
	fmt.Scanln(&player)
}

func gameLogic() {
	computer = generateComputerChoice()
	result = determineWinner(player, computer)
	printGame(player, computer, result)

}
func checkChoice() {
	if player == "rock" || player == "paper" || player == "scissors" {
		gameLogic()
	} else {
		fmt.Println("Invalid choice")
		fmt.Println("Choose rock, paper, or scissors:")
		fmt.Scanln(&player)
		checkChoice()
	}
}
func gameLoop() {
	var choice string
	fmt.Println("Play again? (y/n)")
	fmt.Scanln(&choice)
	if choice == "y" {
		askUserForChoice()
		checkChoice()
		gameLoop()
	} else {
		fmt.Println("Thanks for playing!")
	}
}

func main() {
	askUserForChoice()
	checkChoice()
	gameLoop()
}
