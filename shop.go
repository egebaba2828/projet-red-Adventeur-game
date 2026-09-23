package main

import (
	"fmt"
)

func (c *Character) accessMarchand() {
	for {
		fmt.Println("\n==================================")
		fmt.Println("       BOUTIQUE DU MARCHAND      ")
		fmt.Println("==================================")
		fmt.Printf(" Votre Or : %d pièces\n", c.Gold)
		fmt.Println(" 1. Potion de soin   - 30 or")
		fmt.Println(" 2. Epée en fer (+5) - 80 or")
		fmt.Println(" 3. Quitter la boutique")
		fmt.Println("==================================")
		fmt.Print(" Choisissez une option : ")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			if c.Gold >= 30 && len(c.Inventory) < c.MaxInventory {
				c.Gold -= 30
				c.Inventory = append(c.Inventory, "Potion de soin")
				fmt.Println("\n[Achat réussi] Potion de soin ajoutée à l'inventaire !")
			} else if c.Gold < 30 {
				fmt.Println("\n[Erreur] Pas assez d'or !")
			} else {
				fmt.Println("\n[Erreur] Inventaire plein !")
			}
		} else if choice == 2 {
			if c.Gold >= 80 && len(c.Inventory) < c.MaxInventory {
				c.Gold -= 80
				c.AttackPower += 5
				c.Inventory = append(c.Inventory, "Epée en fer")
				fmt.Println("\n[Achat réussi] Epée en fer équipée ! (+5 Attaque)")
			} else if c.Gold < 80 {
				fmt.Println("\n[Erreur] Pas assez d'or !")
			} else {
				fmt.Println("\n[Erreur] Inventaire plein !")
			}
		} else if choice == 3 {
			fmt.Println("\nAu revoir !")
			break
		} else {
			fmt.Println("\nOption invalide !")
		}
	}
}
