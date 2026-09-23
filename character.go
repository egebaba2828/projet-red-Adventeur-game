package main

import (
	"fmt"
)

type Character struct {
	Name         string
	Class        string
	Level        int
	Exp          int
	MaxExp       int
	MaxHP        int
	CurrentHP    int
	AttackPower  int
	Gold         int
	Inventory    []string
	MaxInventory int
}

func initCharacter(name string, class string, level int, maxHP int, currentHP int, attackPower int, gold int, inventory []string) Character {
	return Character{
		Name:         name,
		Class:        class,
		Level:        level,
		Exp:          0,
		MaxExp:       100,
		MaxHP:        maxHP,
		CurrentHP:    currentHP,
		AttackPower:  attackPower,
		Gold:         gold,
		Inventory:    inventory,
		MaxInventory: 10,
	}
}

func (c *Character) displayInfo() {
	fmt.Println("\n==================================")
	fmt.Println("    INFORMATIONS DU PERSONNAGE   ")
	fmt.Println("==================================")
	fmt.Printf(" Nom        : %s\n", c.Name)
	fmt.Printf(" Classe     : %s\n", c.Class)
	fmt.Printf(" Niveau     : %d (XP: %d/%d)\n", c.Level, c.Exp, c.MaxExp)
	fmt.Printf(" PV         : %d / %d\n", c.CurrentHP, c.MaxHP)
	fmt.Printf(" Attaque    : %d\n", c.AttackPower)
	fmt.Printf(" Or         : %d pièces\n", c.Gold)
	fmt.Printf(" Inventaire : %v (%d/%d)\n", c.Inventory, len(c.Inventory), c.MaxInventory)
	fmt.Println("==================================")
}

func (c *Character) gainExp(amount int) {
	c.Exp += amount
	fmt.Printf("\n[XP] Vous avez gagné %d points d'expérience !\n", amount)

	for c.Exp >= c.MaxExp {
		c.Exp -= c.MaxExp
		c.Level++
		c.MaxExp = int(float64(c.MaxExp) * 1.5)
		c.MaxHP += 20
		c.CurrentHP = c.MaxHP
		c.AttackPower += 5

		fmt.Println("\n==================================")
		fmt.Printf("  LEVEL UP ! Vous êtes niveau %d ! \n", c.Level)
		fmt.Printf("  PV Max: +20 | Attaque: +5\n")
		fmt.Println("==================================")
	}
}

func (c *Character) takePotion() {
	for i, item := range c.Inventory {
		if item == "Potion de soin" {
			c.CurrentHP += 40
			if c.CurrentHP > c.MaxHP {
				c.CurrentHP = c.MaxHP
			}
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			fmt.Printf("\n[Soin] Vous avez utilisé une Potion de soin ! (PV: %d/%d)\n", c.CurrentHP, c.MaxHP)
			return
		}
	}
	fmt.Println("\n[Soin] Vous n'avez pas de Potion de soin dans votre inventaire !")
}
