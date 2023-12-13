package main

import (
	"fmt"
	"github.com/cs50-romain/Match-MakingBazaar/player"
	"github.com/cs50-romain/Match-MakingBazaar/heap"
)

var player_queue = heap.Init()
var player_finishedgame_queue = []*player.Player{}

func fillQueue(players ...*player.Player) {
	for _, player := range players {
		player_queue.Insert(player)
	}
}

func matchMaking() {
	first_player := player_queue.Pop()
	second_player := player_queue.Pop()
	fmt.Printf("Username: %s Skill Level: %d VS Username: %s Skill Level: %d\n", first_player.Username, first_player.Skill_lvl, second_player.Username, second_player.Skill_lvl)

	randomWinner(first_player, second_player)
	player_finishedgame_queue = append(player_finishedgame_queue, first_player, second_player)
}

func randomWinner(player_one, player_two *player.Player) {
	player_one.Skill_lvl++
	player_one.Wins++
	player_two.Skill_lvl--
	player_two.Wins--
	fmt.Printf("%s wins!\n\n", player_one.Username)
}

func printStats(players ...*player.Player) {
	for _, pplayer := range players {
		player.PrintPlayerStat(pplayer)
	}
}

func main() {
	player1 := player.CreatePlayer("gotaga")	
	player2 := player.CreatePlayer("nonane")	
	player3 := player.CreatePlayer("zer0")	
	player4 := player.CreatePlayer("nikof")	
	player5 := player.CreatePlayer("l3tt4ce")	
	player6 := player.CreatePlayer("zenx")	
	player7 := player.CreatePlayer("squeezie")	
	player8 := player.CreatePlayer("broks") 

	fillQueue(player1, player2, player3, player4, player5, player6, player7, player8)

	fmt.Println("Before matches:")
	printStats(player1, player2, player3, player4, player5, player6, player7, player8)
	fmt.Println()

	counter := 0
	for counter < 4 {
		matchMaking()
		counter++
	}

	fmt.Println("After matches:")
	printStats(player1, player2, player3, player4, player5, player6, player7, player8)
}
