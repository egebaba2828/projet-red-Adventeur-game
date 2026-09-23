package main

type Monster struct {
	Name        string
	MaxHP       int
	CurrentHP   int
	AttackPower int
	ExpReward   int
	GoldReward  int
}

func initMonster(name string, maxHP int, attackPower int, expReward int, goldReward int) Monster {
	return Monster{
		Name:        name,
		MaxHP:       maxHP,
		CurrentHP:   maxHP,
		AttackPower: attackPower,
		ExpReward:   expReward,
		GoldReward:  goldReward,
	}
}
