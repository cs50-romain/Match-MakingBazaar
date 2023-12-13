package player

import (
	"fmt"
	"math/rand"
)

type Player struct {
	Username	string
	Skill_lvl	int
	Wins		int
}

func CreatePlayer(username string) *Player {
	return &Player{
		Username: username,
		Skill_lvl: generateRandomSkillLevel(),
		Wins: 0,
	} // Pick a random number for skill level
}

func generateRandomSkillLevel() int {
	return rand.Intn(100 - 1) + 1
}

func PrintPlayerStat(player *Player) {
	fmt.Printf("Username:%s -> |skill level: %d| |Wins:%d|\n", player.Username, player.Skill_lvl, player.Wins)
}
