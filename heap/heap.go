package heap

import "github.com/cs50-romain/Match-MakingBazaar/player"

type Heap struct {
	len	int
	data	[]*player.Player
}

func (h *Heap) Insert(player *player.Player) {
	h.data[h.len] = player
	h.heapifyUp(h.len)
	h.len++
}

func (h *Heap) Pop() *player.Player{
	if h.len == 0 {
		return &player.Player{}
	}

	head_player := h.data[0]
	h.len--
	if h.len == 0 {
		h.data = []*player.Player{}
		return head_player
	}

	h.heapifyDown(0)
	return head_player
}

func (h *Heap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	parent_idx := parentIndex(idx)
	parent_player := h.data[parent_idx]
	player := h.data[idx]

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

	if left_child_idx > h.len {
		return
	}

	if right_child_idx > h.len {
		return
	}

	left_player := h.data[left_child_idx]
	right_player := h.data[right_child_idx]
	player := h.data[idx]

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
	h.data[idx] = value2
	h.data[idx2] = value
}
