package main

import (
	"fmt"
)

func main() {
	player := initCharacter("Aventurier", "Humain", 1, 100, 100, 20, 50, []string{"Potion de soin"})

	for {
		fmt.Println("\n==================================")
		fmt.Println("          MENU PRINCIPAL         ")
		fmt.Println("==================================")
		fmt.Println(" 1. Afficher les informations du personnage")
		fmt.Println(" 2. Accéder à la boutique (Marchand)")
		fmt.Println(" 3. Entrer dans le Donjon (Combattre)")
		fmt.Println(" 4. Quitter le jeu")
		fmt.Println("==================================")
		fmt.Print(" Choisissez une option (1-4) : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			player.displayInfo()
		case 2:
			player.accessMarchand()
		case 3:
			player.chooseDungeon()
		case 4:
			fmt.Println("\nMerci d'avoir joué ! À bientôt.")
			return
		default:
			fmt.Println("\nChoix invalide, veuillez réessayer.")
		}
	}
}
