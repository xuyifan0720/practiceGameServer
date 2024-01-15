package GameRoomHandler

import (
	"github.com/xuyifan0720/rockPaperScissors/GameRoom"
	"net/http"
)

type OperationType int

const (
	Join OperationType = iota
	Remove
)

type arguments struct {
	player int
	room   int
}

type Operation struct {
	ty        OperationType
	args      arguments
	replyChan chan error
}

type GameRoomManager struct {
	room    *GameRoom.GameRoom
	channel chan<- Operation
}

type GameRoomHandler struct {
	managers []GameRoomManager
}

func StartGameRoomHandler(roomCount int) *GameRoomHandler {
	managers := make([]GameRoomManager, 0)
	for i := 0; i < roomCount; i += 1 {
		room := GameRoom.GameRoom{
			PlayerLimit: 2,
			PlayerCount: 0,
			Players:     make([]int, 0),
			State:       GameRoom.Open,
		}
		currentChan := make(chan Operation)
		manager := GameRoomManager{
			room:    &room,
			channel: currentChan,
		}
		managers = append(managers, manager)
		go func() {
			for operation := range currentChan {
				switch operation.ty {
				case Join:
					err := manager.room.JoinPlayer(operation.args.player)
					operation.replyChan <- err
				case Remove:
					err := manager.room.RemovePlayer(operation.args.player)
					operation.replyChan <- err
				}
			}
		}()
	}
	return &GameRoomHandler{managers}
}

func (grh *GameRoomHandler) HandleJoin(w http.ResponseWriter, req *http.Request) {

}
