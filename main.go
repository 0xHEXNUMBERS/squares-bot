package main

import (
	"fmt"

	"github.com/0xhexnumbers/gmcts"
	"github.com/0xhexnumbers/go-squares"
)

const (
	red gmcts.Player = iota
	green
)

var (
	rWin = []gmcts.Player{}
	gWin = []gmcts.Player{}
	draw = []gmcts.Player{red, green}
)

type game struct {
	squares.Game
}

func (g game) GetActions() []gmcts.Action {
	positions := g.Game.GetActions()
	actions := make([]gmcts.Action, len(positions))

	for i, p := range positions {
		actions[i] = p
	}
	return actions
}

func (g game) ApplyAction(a gmcts.Action) (gmcts.Game, error) {
	pos := a.(squares.Position)
	nextState, err := g.Game.ApplyAction(pos)
	return game{nextState}, err
}

func (g game) Player() gmcts.Player {
	if g.Game.Player() == squares.RED {
		return red
	}
	return green
}

func (g game) Winners() []gmcts.Player {
	winner := g.Game.Winner()
	switch winner {
	case squares.RED:
		return rWin
	case squares.GREEN:
		return gWin
	default:
		return draw
	}
}

func main() {
	gameState := game{squares.NewGame()}

	rounds := 10000

	for !gameState.IsTerminal() {
		mcts := gmcts.NewMCTS(gameState)

		tree := mcts.SpawnTree()
		tree.SearchRounds(rounds)
		mcts.AddTree(tree)

		nextState, _ := gameState.ApplyAction(mcts.BestAction())
		gameState = nextState.(game)
		fmt.Println(gameState)
	}
}
