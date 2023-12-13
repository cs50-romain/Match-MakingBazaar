package heap

import (
	"github.com/cs50-romain/Match-MakingBazaar/player"
)

type Heap struct {
	len	int
	Data	[]*player.Player
}

func Init() *Heap {
	return &Heap{0, []*player.Player{}}
}

func (h *Heap) Insert(player *player.Player) {
	if h.len == len(h.Data) {
		h.Data = append(h.Data, player)
	} else {
		h.Data[h.len] = player
	}
	h.heapifyUp(h.len)
	h.len++
}

func (h *Heap) Pop() *player.Player{
	if h.len == 0 {
		return &player.Player{}
	}

	head_player := h.Data[0]
	h.len--
	if h.len == 0 {
		h.Data = []*player.Player{}
		return head_player
	}
	
	h.Data[0] = h.Data[h.len]
	h.heapifyDown(0)
	return head_player
}

func (h *Heap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	parent_idx := parentIndex(idx)
	parent_player := h.Data[parent_idx]
	player := h.Data[idx]

	if player.Skill_lvl > parent_player.Skill_lvl {
		h.swap(player, parent_player, idx, parent_idx)
		h.heapifyUp(parent_idx)
	}
}

func (h *Heap) heapifyDown(idx int) {
	if idx == h.len {
		return
	}

	left_child_idx := leftChildIndex(idx)
	right_child_idx := rightChildIndex(idx)

	if left_child_idx >= h.len {
		return
	}

	if right_child_idx >= h.len {
		return
	}

	left_player := h.Data[left_child_idx]
	right_player := h.Data[right_child_idx]
	player := h.Data[idx]

	if right_player.Skill_lvl > left_player.Skill_lvl && player.Skill_lvl < right_player.Skill_lvl {
		h.swap(player, right_player, idx, right_child_idx)
		h.heapifyDown(right_child_idx)
	} else if left_player.Skill_lvl > right_player.Skill_lvl && player.Skill_lvl < left_player.Skill_lvl {
		h.swap(player, left_player, idx, left_child_idx)	
	}
}

func parentIndex(idx int) int {
	return (idx - 1) / 2
}

func leftChildIndex(idx int) int {
	return 2 * idx + 1
}

func rightChildIndex(idx int) int {
	return 2 * idx + 2
}

func (h *Heap) swap(value, value2 *player.Player, idx, idx2 int) {
	h.Data[idx] = value2
	h.Data[idx2] = value
}
