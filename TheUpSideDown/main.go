package main

import "fmt"

func main() {

	fmt.Println("The Voice: Hello Traveler, Welcome to the UpsideDown.")

	var name string
	fmt.Println("The Voice: Who are you?")
	fmt.Printf("Enter Your Name:\t")

	fmt.Scan(&name)

	fmt.Println("\nThe Voice: DOOM AWAITS!! Enter if you dare 💀 💀")

	var choice string

	fmt.Printf("Continue?:\t")
	fmt.Scan(&choice)

	if choice == "yes" {
		fmt.Println("Hahahaha...")
	} else {
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

	var weapon string
	fmt.Printf("\nChoose a weapon:\n")

	weapon1 := "Axe"
	fmt.Println(weapon1, " 🪓")

	weapon2 := "Sword"
	fmt.Println(weapon2, " 🗡️")

	weapon3 := "Arrow"
	fmt.Println(weapon3, " 🏹")

	fmt.Printf("\n")
	fmt.Scan(&weapon)

	if weapon == "Axe" {
		fmt.Println("Player selected: your weapon is an 🪓")
	} else if weapon == "Sword" {
		fmt.Println("Player selected: your weapon is a Sword 🗡️")
	} else if weapon == "Arrow" {
		fmt.Println("Player selected: your weapon is an Arrow 🏹")
	} else {
		return
	}

	fmt.Println("\n----------------------------------    LEVEL 1    -----------------------------------------------")

	fmt.Println("You're in the DemiGorgon's lair. You are attacked by one of his minions. ")
	fmt.Println("\n                        THE KRAKEN 🐙   [🖤🖤🖤🖤🖤]     ")

	var lives string
	fmt.Println(lives, "\nYour Health: [ ❤️ ❤️ ❤️ ]")

	fmt.Println("\nThe Voice: The Kraken attacks from the left.")

	var move string

	fmt.Printf("Enter:\t")
	fmt.Scan(&move)

	if move == "A" {
		fmt.Println("\nAttack evaded")
	} else if move == "a" {
		fmt.Println("\nAttack evaded")
	} else if move != "A" {
		fmt.Println("\nYou Got Hit! [ ❤️ ❤️ ]")
	} else if move != "a" {
		fmt.Println("\nYou Got Hit! [ ❤️ ❤️ ]")
	}

	var attack string

	fmt.Println("YOUR TURN")
	fmt.Println("\nStrike: S")
	fmt.Println("Punch: P")

	fmt.Printf("\nEnter:\t")
	fmt.Scan(&attack)

	if attack == "S" {
		fmt.Println("\nYou chopped off a tentacle. 🐙: [🖤🖤🖤🖤]")
	} else if attack == "P" {
		fmt.Println("\n🐙: You Missed Me")

		fmt.Println("\nThe Voice: Use Melee to deal heavy damage to enemies")
	}

	fmt.Println("\nThe Kraken opens its mouth and spits out Octo-goo 💦")

	var move1 string

	fmt.Println("\nMake a move:")
	fmt.Println("Evade: E")
	fmt.Println("Strike: S")
	fmt.Println("\n ")

	fmt.Printf("Enter:\t")
	fmt.Scan(&move1)

	if move1 == "E" {
		fmt.Println("\nAttack evaded")
	} else if move1 == "e" {
		fmt.Println("\nAttack evaded")
	} else if move1 == "S" {
		fmt.Println("\nYou Got Hit! [ ❤️ ]")
	} else if move1 == "s" {
		fmt.Println("\nYou Got Hit! [ ❤️ ]")
	}

	



}
