package player

import (
	"fmt"
	"math/rand"
	"practice/snakes-and-ladders/board"
)

type PlayerDetails struct {
	name     string
	isTurn   bool
	isWinner bool
	pos      int
	board    board.BoardGame
}

type Player interface {
	Play(n int)
	RollDice() int
	SetIsTurn(b bool)
	IsWinner() bool
	GetName() string
}

func (p *PlayerDetails) RollDice() int {
	return 1 + rand.Intn(6)
}
func (p *PlayerDetails) SetIsTurn(b bool) {
	p.isTurn = b
}
func (p *PlayerDetails) IsWinner() bool {
	return p.isWinner
}
func (p *PlayerDetails) GetName() string {
	return p.name
}
func (p *PlayerDetails) Play(n int) {
	// if got 6, player will play again
	jump := p.RollDice()
	fmt.Printf("%v rolled to %v\n", p.name, jump)
	newPos := p.pos + jump
	if newPos >= n*n {
		return
	}
	row := (newPos / 10)
	col := (newPos % 10)
	spot, err := p.board.GetSpot(row, col)
	if err != nil {
		fmt.Printf("Error A:%v,row:%v,col:%v,pos:%v\n", err, row, col, newPos)
		return
	}
	p.pos = spot.GetPos()
	spotType := spot.GetPosType()
	nextSpot := spot.GetNextPos()
	if spotType != "normal" {
		newPos := nextSpot
		fmt.Printf("%v got to %v from %v by:%v\n", p.name, newPos, p.pos, spotType)
		row := (newPos / 10)
		col := (newPos % 10)
		spot, err := p.board.GetSpot(row, col)
		if err != nil {
			fmt.Printf("Error B:%v\n", err)
			return
		}
		p.pos = spot.GetPos()

	}
	p.isTurn = false
	if p.pos == (n*n)-1 {
		p.isWinner = true
	}
	fmt.Printf("%v moves to %v\n", p.name, p.pos)
	return

}

func CreateNewPlayer(name string, board board.BoardGame) Player {
	return &PlayerDetails{
		name:  name,
		board: board,
		pos:   0,
	}
}
