package board

import (
	"errors"
	"practice/snakes-and-ladders/spot"
)

type Board struct {
	board [][]spot.Spot
	size  int
}

type BoardGame interface {
	GetBoard() [][]spot.Spot
	GetSpot(i, j int) (spot.Spot, error)
	AddNextPos(i, j, nextPos int, posType string) (error, bool)
}

func (b *Board) GetBoard() [][]spot.Spot {
	return b.board
}

func (b *Board) GetSpot(i, j int) (spot.Spot, error) {
	//validate i and j
	if i < b.size && j < b.size {
		return b.board[i][j], nil
	}
	// fmt.Printf("Error", i, j, b.size)
	return &spot.SpotData{}, errors.New("Invalid")

}

func (b *Board) AddNextPos(i, j, nextPos int, posType string) (error, bool) {
	if i < b.size && j < b.size {
		spot := b.board[i][j]
		ans := spot.SetNextPos(nextPos, posType)
		return nil, ans
	} else {
		return errors.New("Error"), false
	}

}

func CreateNewBoard(n int) BoardGame {
	board := make([][]spot.Spot, n)
	for i := range board {
		board[i] = make([]spot.Spot, n)
	}
	k := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = spot.CreateNewSpot(k, k+1, "normal")
			k++
		}
	}
	return &Board{
		board: board,
		size:  n,
	}
}
