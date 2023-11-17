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

	fmt.Println("\n- Do You Wish To Continue?\t")
	fmt.Println("\nYes: y")
	fmt.Println("No: n")

	fmt.Printf("\nEnter:\t")
	fmt.Scan(&choice)

	if choice == "y" {
		fmt.Println("Hahahaha...")
	} else if choice == "n" {
		return
	} else {
		fmt.Println("Invalid Entry")
	}

	if choice != "y" {
		return
	}

	fmt.Println("\nSolve the riddle: Hey baby, if you've got this methods then, you're my type")

	var answer string = "interface"

	fmt.Printf("Enter your answer:\t")
	fmt.Scan(&answer)

	if answer == "interface" {
		fmt.Printf("\nChoose A Weapon:\n")
	} else {
		fmt.Println("You Failed the Riddle!")
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
	} else if weapon == "axe" {
		fmt.Println("Weapon selected: your weapon is an Axe 🪓")
	} else if weapon == "Sword" {
		fmt.Println("Weapon selected: your weapon is a Sword 🗡️")
	} else if weapon == "sword" {
		fmt.Println("Weapon selected: your weapon is a Sword 🗡️")
	} else if weapon == "Arrow" {
		fmt.Println("Weapon selected: your weapon is an Arrow 🏹")
	} else if weapon == "arrow" {
		fmt.Println("Weapon selected: your weapon is an Arrow 🏹")
	} else {
		return
	}

	fmt.Println("\n------------------------    START   ----------------------------")

	fmt.Println("\n- You're in the DemiGorgon's lair. You are attacked by one of his minions. ")
	fmt.Println("\n\n                        THE KRAKEN 🐙   [🖤🖤🖤🖤]     ")

	var lives string
	fmt.Println(lives, "\nYour Health: [ ❤️ ]")

	fmt.Println("\n- The Kraken attacks from the left. Make a move")

	var move string

	fmt.Println("\nEvade: E")
	fmt.Println("Strike: S")
	fmt.Println("Block: B")
	fmt.Println("\n ")

	fmt.Printf("Enter:\t")
	fmt.Scan(&move)

	if move == "e" {
		fmt.Println("You Evaded The Attack")
	} else if move == "s" {
		fmt.Println("\nThe Kraken threw a surprise attack, You Got Hit! Game over: [ 💔 ] ")
	} else if move == "b" {
		fmt.Println("\nThe Kraken broke your defense, You Got Hit! Game over: [ 💔 ] ")
	} else {
		fmt.Println("Invalid Entry")
	}

	// End The Game If Player Does not Evade the Enemy Attack
	if move != "e" {
		return
	}

	var attack string

	fmt.Println("\n- The Kraken still in full health looses its guard. You see an opening")

	fmt.Println("\nStrike: S")
	fmt.Println("Punch: P")

	fmt.Printf("\nEnter:\t")
	fmt.Scan(&attack)

	if attack == "s" {
		fmt.Println("\nYou chopped off a tentacle. 🐙: [🖤🖤🖤]")
	} else if attack == "p" {
		fmt.Println("\n🐙: You Missed Me")
		fmt.Println("\nYou Got Hit! Game over: [ 💔 ] ")
	} else {
		fmt.Println("Invalid Entry")
	}

	// End The Game If Player Makes a Wrong Choice
	if attack != "s" {
		return
	}

	fmt.Println("\nThe Kraken opens its mouth and spits out Octo-goo 💦")

	var move1 string

	fmt.Println("\nMake a move:")

	fmt.Println("Evade: E")
	fmt.Println("Strike: S")
	fmt.Println("Block: B")

	fmt.Println("\n ")

	fmt.Printf("Enter:\t")
	fmt.Scan(&move1)

	if move1 == "e" {
		fmt.Println("You Evaded The Attack")
	} else if move1 == "s" {
		fmt.Println("\nYou Got Hit! Game over: [ 💔 ] ")
	} else if move1 == "b" {
		fmt.Println("\n🐙: Your defenses against me are weak!")
		fmt.Println("Got Hit! Game over: [ 💔 ] ")
	} else {
		fmt.Println("Invalid Entry")
	}

	// End The Game If Player Does not Evade the Enemy Attack
	if move1 != "e" {
		return
	}

	var attack1 string

	fmt.Println("\n- The Kraken Launches at you from the Top")

	fmt.Println("\nStrike: S")
	fmt.Println("Parry: P")
	fmt.Println("Block: B")
	fmt.Println("Evade: E")

	fmt.Printf("\nEnter:\t")
	fmt.Scan(&attack1)

	if attack1 == "s" {
		fmt.Println("\nToo Slow. You Got Hit! Game over: [ 💔 ] ")
	} else if attack1 == "p" {
		fmt.Println("\nYou Deflected its attack you see an opening")
	} else if attack1 == "b" {
		fmt.Println("You Blocked its attack")
	} else if attack1 == "e" {
		fmt.Println("You Evaded The Attack")
	} else {
		fmt.Println("Invalid Entry")
	}

	if attack1 == "s" {
		return
	}

	var chance string

	if attack1 == "p" {
		fmt.Println("Strike: S")
		fmt.Printf("\nEnter:\t")
		fmt.Scan(&chance)

		fmt.Println("\nYou chopped off a tentacle. 🐙: [🖤🖤]")
		fmt.Println("\nThe Kraken opens its mouth and spits out Octo-goo 💦")
	}

	if attack1 == "b" {
		fmt.Println("\n🐙: Arrrghh!!!!!!!!!!!")
		fmt.Println("\nThe Kraken opens its mouth and spits out Octo-goo 💦")
	}

	var chance0 string

	fmt.Println("\nMake a move:")
	fmt.Println("Evade: E")
	fmt.Println("Strike: S")
	fmt.Println("Block: B")

	fmt.Printf("\nEnter:\t")
	fmt.Scan(&chance0)

	if chance0 == "e" {
		fmt.Println("You Evaded The Attack")
	} else if chance0 == "s" {
		fmt.Println("🐙: Can't fool me twice... You Got Hit! Game over: [ 💔 ] ")
	} else if chance0 == "b" {
		fmt.Println("Not So Brave. Game over: [ 💔 ] ")
	} else {
		fmt.Println("Invalid Entry")
	}

	if chance0 != "e" {
		return
	}

	fmt.Println("\nThe Walls Of The Caves Collaps, You See An Exit And Try To Escape...")
	fmt.Println("But Your Way Is Blocked By The Kraken")

	fmt.Println("\nThe Kraken Launches A Series Of Attack. Left, Left, Right")

	var evade string

	fmt.Println("\nMake a move:")
	fmt.Println("⬅️ ⬅️ ➡️: LLR")
	fmt.Println("➡️ ➡️ ⬅️: RRL")

	fmt.Printf("\nEnter:\t")
	fmt.Scan(&evade)

	if evade == "llr" {
		fmt.Println("\nYou Dodged The Attack Stunning The Kraken With A Heavy Blow Killing The Kraken. 🐙: [ 💀 ]")
	} else if evade == "rrl" {
		fmt.Println("\nHappened So Fast... Game over: [ 💔 ]")
	}

	if evade != "llr" {
		return
	}

	fmt.Println("\n\t\t\t\t\tTHE END")

	fmt.Println("\n\t\t\tThank You For Playing My Game.")

}
