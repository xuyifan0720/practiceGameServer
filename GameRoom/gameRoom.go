package GameRoom

import (
	"github.com/xuyifan0720/rockPaperScissors/Utils"
)

type Roomstate int

const (
	Open Roomstate = iota
	InProgress
	Finished
)

type GameRoom struct {
	PlayerLimit int
	PlayerCount int
	Players     []int
	State       Roomstate
}

type roomOperationError struct {
	message string
}

func (r *roomOperationError) Error() string {
	return r.message
}

func (r *GameRoom) JoinPlayer(player int) error {
	if r.PlayerCount >= r.PlayerLimit {
		return &roomOperationError{"room is already full"}
	}
	idx := Utils.FindElements(r.Players, player)
	if idx != -1 {
		return &roomOperationError{"Player is already in the room"}
	}
	r.PlayerCount += 1
	r.Players = append(r.Players, player)
	return nil
}

func (r *GameRoom) RemovePlayer(player int) error {
	if r.PlayerCount == 0 {
		return &roomOperationError{"room is empty"}
	}
	idx := Utils.FindElements(r.Players, player)
	if idx == -1 {
		return &roomOperationError{"Player is not in the room"}
	}
	r.PlayerCount -= 1
	r.Players = append(r.Players[:idx], r.Players[idx+1:]...)
	return nil
}
