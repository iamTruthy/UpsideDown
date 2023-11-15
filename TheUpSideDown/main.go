package main

import "fmt"

func main() {

	fmt.Println("The Voice: Hello Traveler, Welcome to the UpsideDown.")

	var name string
	fmt.Println("The Voice: Who are you?")
	fmt.Printf("Enter your name:\t")

	fmt.Scan(&name)

	fmt.Println("\nThe Voice: DOOM AWAITS!! Enter if you dare 💀 💀")

	var choice string

	fmt.Printf("Continue?:\t")
	fmt.Scan(&choice)

	if choice == "yes" {
		fmt.Println("Hahahaha...")
	} else if choice == "no" {
		return
	}

	fmt.Println("\nSolve the riddle:Hey baby, if you've got this methods then, you're my type")
	fmt.Printf("Enter your answer:\t")

	var answer string = "interface"

	if answer == "interface" {
		fmt.Scan(&answer)
	} else {
		return
	}

	var player string
	fmt.Printf("\nChoose a player:\n")

	player1 := "Tweleven"
	fmt.Println(player1)

	player2 := "Steven"
	fmt.Println(player2)

	player3 := "Bustin"
	fmt.Println(player3)

	fmt.Printf("\n")
	fmt.Scan(&player)

	if player == "Tweleven" {
		fmt.Println("Player selected: your weapon is the FORCE ✋")
	} else if player == "Steven" {
		fmt.Println("Player selected: your weapon is a Batton 🏏")
	} else if player == "Bustin" {
		fmt.Println("Player selected: your weapon is a Ray gun 🔫")
	} else {
		return
	}

	fmt.Println("\n----------------------------------    LEVEL 1    -----------------------------------------------")

	fmt.Println("You're in the DemiGorgon's lair. You are attacked by one of his minions. ")
	fmt.Println("\n                       THE KRAKEN 🐙        ")

	lives := "Lives"
	fmt.Println(lives, ":❤️ ❤️ ❤️")

	fmt.Println("\nThe Voice: The Kraken attacks from the left.")

	var move string

	fmt.Printf("Your Move:\t")
	fmt.Scan(&move)

	if move == "A" {
		fmt.Println("Attack evaded")
	} else if move == "a" {
		fmt.Println("Attack evaded")
	} else if move != "A" {
		fmt.Println("You Got Hit!")
	} else if move != "a" {
		fmt.Println("You Got Hit")
	}

	var attack string

	fmt.Printf("\nYour Turn:\t")
	fmt.Scan(&attack)

	if attack == "Shoot" {
		fmt.Println("Kraken: 🖤 🖤 🖤")
	} else if attack == "Punch" {
		fmt.Println("You missed")
	}
}
