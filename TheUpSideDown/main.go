package main

import "fmt"

func main() {

	fmt.Println("The Voice: Hello Traveler, Welcome to the UpsideDown.")

	var name string

	fmt.Printf("The Voice: Who are you?\n")

	fmt.Scan(&name)

	fmt.Println("The Voice: Enter if you dare!!")

	var choice string

	fmt.Printf("Continue?\n")
	fmt.Scan(&choice)

	if choice == "yes" {
		fmt.Println("\nDOOM AWAITS!!!...")
	} else if choice == "no" {
		fmt.Println("Hahahaha...")
		return
	}

	var player string

	fmt.Printf("Choose a player\n")

	p1 := "Tweleven"
	fmt.Println("\n", p1)

	p2 := "Steven"
	fmt.Println("\n", p2)

	p3 := "Bustin"
	fmt.Println("\n", p3)

	fmt.Printf("\n")

	fmt.Scan(&player)

	if player == "Tweleven" {
		fmt.Println("Player selected: your weapon is the FORCE")
	} else if player == "Steven" {
		fmt.Println("Player selected: your weapon is Lucile")
	} else if player == "Bustin" {
		fmt.Println("Player selected: your weapon is a Ray gun")
	} else {
		return
	}

	fmt.Println("\n--------------------------  LEVEL 1  ------------------------------------------")

	fmt.Println("You're in DemiGorgon's lair. You are attacked by one of his minions. A boulder is thrown your way.")

}
