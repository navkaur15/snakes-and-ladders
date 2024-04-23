package main

import (
	"fmt"
	"practice/snakes-and-ladders/board"
	"practice/snakes-and-ladders/player"
)

func main() {
	n := 10
	board := board.CreateNewBoard(n)
	board.AddNextPos(1, 7, 7, "snake")
	board.AddNextPos(5, 4, 34, "snake")
	board.AddNextPos(6, 2, 81, "ladder")
	board.AddNextPos(8, 7, 96, "ladder")
	board.AddNextPos(1, 3, 67, "ladder")

	//set next pos (for snakes and ladders)
	p1 := player.CreateNewPlayer("Jack", board)
	p2 := player.CreateNewPlayer("John", board)
	players := []player.Player{p1, p2}
	//Play
	flag := false
	for {
		for _, p := range players {
			p.SetIsTurn(true)
			p.Play(n)
			if p.IsWinner() {
				fmt.Printf("%v is winner\n", p.GetName())
				flag = true
				break
			}
		}
		if flag {
			break
		}

	}
	fmt.Println("End")

}
