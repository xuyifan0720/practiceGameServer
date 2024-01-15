package main

import (
	"fmt"

	"github.com/xuyifan0720/rockPaperScissors/GameRoom"
	"github.com/xuyifan0720/rockPaperScissors/GameRoomHandler"
)

func main() {
	fmt.Println("hello world")
	room := &GameRoom.GameRoom{
		2,
		0,
		make([]int, 0),
		GameRoom.Open,
	}
	err := room.JoinPlayer(1)
	fmt.Println(err)
	err = room.JoinPlayer(1)
	fmt.Println(err)
	err = room.JoinPlayer(2)
	fmt.Println(err)
	err = room.JoinPlayer(3)
	fmt.Println(err)
	fmt.Println(room.Players)
	err = room.RemovePlayer(2)
	fmt.Println(err)
	fmt.Println(room.Players)
	err = room.RemovePlayer(2)
	fmt.Println(err)
	err = room.RemovePlayer(1)
	fmt.Println(err)
	err = room.RemovePlayer(2)
	fmt.Println(err)
	handler := GameRoomHandler.StartGameRoomHandler(3)
	fmt.Println(handler)
}
