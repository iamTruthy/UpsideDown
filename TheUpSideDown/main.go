package main

import "fmt"

func main() {

	fmt.Println("🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱🚪 THE UPSIDEDOWN 🚪🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱")

	fmt.Println("\nThe Voice: Hello Traveler, Welcome to the UpsideDown.")

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
	} else if choice == "no" {
		return
	} else {
		fmt.Println("Invalid Entry")
	}

	if choice != "yes" {
		return
	}

	fmt.Println("\nSolve the riddle:Hey baby, if you've got this methods then, you're my type")

	var answer string = "interface"

	fmt.Printf("Enter your answer:\t")
	fmt.Scan(&answer)

	if answer == "interface" {
		fmt.Printf("\nChoose A Weapon:\n")
	} else if answer != "interface" {
		fmt.Println("Invalid Entry")
	}

	if answer != "interface" {
		return
	}

	var weapon string

	weapon1 := "Axe"
	fmt.Println(weapon1, " 🪓")

	weapon2 := "Sword"
	fmt.Println(weapon2, " 🗡️")

	weapon3 := "Arrow"
	fmt.Println(weapon3, " 🏹")

	fmt.Printf("\nEnter Weapon:\t")
	fmt.Scan(&weapon)

	if weapon == "Axe" {
		fmt.Println("Weapon selected: your weapon is an Axe 🪓")
	} else if weapon == "Sword" {
		fmt.Println("Weapon selected: your weapon is a Sword 🗡️")
	} else if weapon == "Arrow" {
		fmt.Println("Weapon selected: your weapon is an Arrow 🏹")
	} else {
		return
	}

	fmt.Println("\n----------------------------------    START   -----------------------------------------------")

	fmt.Println("You're in the DemiGorgon's lair. You are attacked by one of his minions. ")
	fmt.Println("\n                        THE KRAKEN 🐙   [🖤🖤🖤🖤]     ")

	var lives string
	fmt.Println(lives, "\nYour Health: [ ❤️ ]")

	fmt.Println("\nThe Voice: The Kraken attacks from the left. Make a move")

	var move string

	fmt.Println("\nEvade: E")
	fmt.Println("Strike: S")
	fmt.Println("\n ")

	fmt.Printf("Enter:\t")
	fmt.Scan(&move)

	if move == "E" {
		fmt.Println("\nThe Voice: Attack evaded")
	} else if move == "e" {
		fmt.Println("\nThe Voice: Attack evaded")
	} else if move != "S" {
		fmt.Println("\nThe Kraken threw a surprise attack, You Got Hit! Game over: [ 💔 ] ")
	} else if move != "s" {
		fmt.Println("\nThe Kraken threw a surprise attack, You Got Hit! Game over: [ 💔 ] ")
	} else {
		fmt.Println("Invalid Entry")
	}

	// End The Game If Player Does not Evade the Enemy Attack
	if move != "E" {
		return
	}

	var attack string

	fmt.Println("\nYOUR TURN")
	fmt.Println("The Kraken still in full health looses its guard. You see an opening")

	fmt.Println("\nStrike: S")
	fmt.Println("Punch: P")

	fmt.Printf("\nEnter:\t")
	fmt.Scan(&attack)

	if attack == "S" {
		fmt.Println("\nYou chopped off a tentacle. 🐙: [🖤🖤🖤]")
	} else if attack == "P" {
		fmt.Println("\n🐙: You Missed Me")
		fmt.Println("\nYou Got Hit! Game over: [ 💔 ] ")
	} else {
		fmt.Println("Invalid Entry")
	}

	// End The Game If Player Makes a Wrong Choice
	if attack != "S" {
		return
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
		fmt.Println("\nYou Got Hit! Game over: [ 💔 ] ")
	} else if move1 == "s" {
		fmt.Println("\nYou Got Hit! Game over: [ 💔 ] ")
	} else {
		fmt.Println("Invalid Entry")
	}

	// End The Game If Player Does not Evade the Enemy Attack
	if move1 != "E" {
		return
	}

	var attack1 string

	fmt.Println("\nThe Kraken Launches at you from the Top")

	fmt.Println("\nStrike: S")
	fmt.Println("Parry: P")
	fmt.Println("Block: B")
	fmt.Println("Evade: E")

	fmt.Printf("\nEnter:\t")
	fmt.Scan(&attack1)

	if attack1 == "S" {
		fmt.Println("\nYou Got Hit! Game over: [ 💔 ] ")
	} else if attack1 == "P" {
		fmt.Println("\nYou Deflected its attack you see an opening")
	} else if attack1 == "B" {
		fmt.Println("You Blocked its attack")
	} else if attack1 == "E" {
		fmt.Println("Attack Evaded")
	} else {
		fmt.Println("Invalid Entry")
	}

	if attack1 == "S" {
		return
	}

	var chance string

	if attack1 == "P" {
		fmt.Println("Strike: S")
		fmt.Printf("\nEnter:\t")
		fmt.Scan(&chance)

		fmt.Println("\nYou chopped off a tentacle. 🐙: [🖤🖤]")
	}

	if attack1 == "B" {
		fmt.Println()
	}

}
