package main

import (
	. "math"
)

const MINVAL = -(1e10)

// constant OppMove = Computer Move, UsrMove = Move Made by Usr
const OppMove = 'X'
const UsrMove = '0'

// returns a slice of all possible moves that can be applied at given time
func _possibleMoves(Grid string) []int {
	var moves []int
	for i := 0; i < 9; i++ {
		if Grid[i] == '.' {
			moves = append(moves, i)
		}
	}
	return moves
}

// helper function to fill the value of grid
func _replace(Grid string, pos int, MOVE rune) string {
	return Grid[:pos] + string(MOVE) + Grid[pos+1:]
}

// returns true if the game is finished else returns false
func checkFinish(Grid string) bool {
	var checkHorizontal bool = false
	var checkVertical bool = false
	for i := 0; i < 3; i++ {
		h := 3 * i
		checkHorizontal = checkHorizontal || (Grid[h] == Grid[h+1] && Grid[h+1] == Grid[h+2] && Grid[h+1] != '.')
		checkVertical = checkVertical || (Grid[i] == Grid[i+3] && Grid[i+3] == Grid[i+6] && Grid[i+3] != '.')
	}
	var checkDiagonal bool = ((Grid[0] == Grid[4] && Grid[4] == Grid[8] && Grid[4] != '.') || (Grid[2] == Grid[4] && Grid[4] == Grid[6] && Grid[4] != '.'))
	return checkDiagonal || checkHorizontal || checkVertical
}

// Optimality of a move = # of solutions that can be generated from taking that move as the next immediate move
// Optimality of a move decreases as we go down the tree of available moves that lead to the solution
func GetOpt(Grid string, player rune, depth int) (opt float64) {
	var factor float64 = (Pow(float64(10), float64(depth)))
	if checkFinish(Grid) {
		if player == OppMove {
			return -1 / factor
		} else {
			return 1 / factor
		}
	}
	var moves []int = _possibleMoves(Grid)
	if len(moves) == 0 {
		return 0
	}
	opt = 0
	var nextp rune
	if player == UsrMove {
		nextp = OppMove
	} else {
		nextp = UsrMove
	}
	n := len(moves)
	for i := 0; i < n; i++ {
		Grid = _replace(Grid, moves[i], player)
		opt += GetOpt(Grid, nextp, depth+1)
		Grid = _replace(Grid, moves[i], rune('.'))
	}
	return opt
}

// Function which provides the next most optimal move for the Computer
func GetOptMove(Grid string) (pos int8) {
	var moves []int = _possibleMoves(Grid)
	pos = -2
	var cval, val float64
	val = MINVAL
	for i := 0; i < len(moves); i++ {
		Grid = _replace(Grid, moves[i], OppMove)
		cval = GetOpt(Grid, UsrMove, 0)
		Grid = _replace(Grid, moves[i], '.')
		if cval > val {
			pos = int8(moves[i])
			val = cval
		}
	}
	return pos + 1
}
