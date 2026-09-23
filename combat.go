package main

import (
	"fmt"
	"time"
)

func (c *Character) startCombat(m *Monster) {
	fmt.Printf("\n==================================\n")
	fmt.Printf("   DEBUT DU COMBAT : %s VS %s   \n", c.Name, m.Name)
	fmt.Printf("==================================\n")

	tour := 1

	for c.CurrentHP > 0 && m.CurrentHP > 0 {
		fmt.Printf("\n--- Tour %d ---\n", tour)
		fmt.Printf("%s (PV: %d/%d) | %s (PV: %d/%d)\n", c.Name, c.CurrentHP, c.MaxHP, m.Name, m.CurrentHP, m.MaxHP)
		fmt.Println("1. Attaquer")
		fmt.Println("2. Utiliser une Potion de soin")
		fmt.Print("Votre choix : ")

		var action int
		fmt.Scan(&action)

		if action == 1 {
			fmt.Printf("\n-> Vous attaquez le %s et infligez %d dégâts !\n", m.Name, c.AttackPower)
			m.CurrentHP -= c.AttackPower
			if m.CurrentHP < 0 {
				m.CurrentHP = 0
			}
		} else if action == 2 {
			c.takePotion()
		} else {
			fmt.Println("\nOption invalide ! Vous perdez votre tour !")
		}

		if m.CurrentHP == 0 {
			fmt.Printf("\nVICTOIRE ! Vous avez vaincu le %s !\n", m.Name)
			c.Gold += m.GoldReward
			fmt.Printf("Vous gagnez %d pièces d'or !\n", m.GoldReward)
			c.gainExp(m.ExpReward)
			break
		}

		time.Sleep(300 * time.Millisecond)

		fmt.Printf("<- Le %s vous attaque et inflige %d dégâts !\n", m.Name, m.AttackPower)
		c.CurrentHP -= m.AttackPower
		if c.CurrentHP < 0 {
			c.CurrentHP = 0
		}

		if c.CurrentHP == 0 {
			fmt.Printf("\nDEFAITE... Vous avez été vaincu par le %s.\n", m.Name)
			break
		}

		tour++
	}
}

func (c *Character) chooseDungeon() {
	if c.CurrentHP <= 0 {
		fmt.Println("\nVous êtes inconscient ! Soignez-vous avant de combattre.")
		return
	}

	fmt.Println("\n==================================")
	fmt.Println("        CHOIX DU DONJON           ")
	fmt.Println("==================================")
	fmt.Println(" 1. Gobelin (Facile)   - XP: 40 | Or: 30")
	fmt.Println(" 2. Orc (Moyen)        - XP: 80 | Or: 60")
	fmt.Println(" 3. Dragon (Difficile) - XP: 200 | Or: 150")
	fmt.Println("==================================")
	fmt.Print(" Choisissez un monstre : ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		goblin := initMonster("Gobelin", 40, 8, 40, 30)
		c.startCombat(&goblin)
	case 2:
		orc := initMonster("Orc Guerrier", 80, 16, 80, 60)
		c.startCombat(&orc)
	case 3:
		dragon := initMonster("Dragon Rouge", 160, 25, 200, 150)
		c.startCombat(&dragon)
	default:
		fmt.Println("\nChoix invalide !")
	}
}
