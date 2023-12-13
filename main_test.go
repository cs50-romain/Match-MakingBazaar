package main

import (
	"testing"
	"github.com/cs50-romain/Match-MakingBazaar/player"
	"github.com/cs50-romain/Match-MakingBazaar/heap"
)

func TestFillQueue(t *testing.T) {
	heap := heap.Init()

	player1 := &player.Player{"gotaga", 99, 0}
	player2 := &player.Player{"nonane", 76, 0}
	player3 := &player.Player{"zer0", 98, 0}
	player4 := &player.Player{"nikof", 24, 0}

	heap.Insert(player1)
	heap.Insert(player2)
	heap.Insert(player3)
	heap.Insert(player4)

	if heap.Data[0].Skill_lvl != 99 {
		t.Errorf("Heap insertion was incorrect, got: %d, want: %d.", heap.Data[0].Skill_lvl, 99)
	}

	if heap.Data[1].Skill_lvl != 76 {
		t.Errorf("Heap insertion was incorrect, got: %d, want: %d.", heap.Data[1].Skill_lvl, 76)
	}

	if heap.Data[2].Skill_lvl != 98 {
		t.Errorf("Heap insertion was incorrect, got: %d, want: %d.", heap.Data[2].Skill_lvl, 98)
	}

	if heap.Data[3].Skill_lvl != 24 {
		t.Errorf("Heap insertion was incorrect, got: %d, want: %d.", heap.Data[3].Skill_lvl, 24)
	}
}
