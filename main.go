package main

import (
	"fmt"
	"log"
	"os"
)

const SizeOfTetromino = 4

type tetromino [SizeOfTetromino][SizeOfTetromino]rune

func main() {
	// Checking input
	if len(os.Args) != 2 {
		fmt.Println("Usage:\ngo run . <filename>\nExample:\ngo run . goodexample00.txt")
		return
	}

	// Read and edit tetrominoes
	tetrominoes, err := inputTetrominoes()
	if err!=nil {
		log.Fatalf("ERROR: %v", err)
		return
	}

	squareSize := 2
	square := makeSquare(squareSize)
	// try to put the tetrominoes in the square
	for !isAssemble(tetrominoes, square) {
		squareSize++
		square = makeSquare(squareSize)
	}

	// Print result
	dotsCounter:=0
	for i := 0; i < len(square); i++ {
		for j := 0; j < len(square); j++ {
			fmt.Print(string(square[i][j]))
			if square[i][j]=='.' {dotsCounter++}
		}
		fmt.Println()
	}
	fmt.Printf("\nthere are %d dots in the square\n", dotsCounter)
}

/*
creates an "empty" square (filled with dots)
*/
func makeSquare(size int) [][]rune {
	square := [][]rune{}
	for i := 0; i < size; i++ {
		square = append(square, []rune{})
		for j := 0; j < size; j++ {
			square[i] = append(square[i], '.')
		}
	}
	return square
}

/*
 */
func isAssemble(tetrominoes []tetromino, square [][]rune) bool {
	if len(tetrominoes) == 0 {
		return true
	}

	// try to place the tetrominoes into the square
	for i := 0; i < len(square); i++ {
		for j := 0; j < len(square); j++ {
			if square[i][j] == '.' { // blank spot in the square
				if putTetrominoInSquare(tetrominoes[0], square, i, j) {
					if isAssemble(tetrominoes[1:], square) {
						return true
					}
					// if it was impossible to place the next tetrominoes delete the current one
					deleteTetrominoFromSquare(tetrominoes[0], square, i, j)
				}
			}
		}
	}
	// cannot place all the tetrominoes into the square
	return false // If tetrominoes don't fit in given square, return false
}

/*
checks if the tetromino fits to the given place into the square
*/
func putTetrominoInSquare(t tetromino, square [][]rune, x, y int) bool {
	piecesOfTetromino := 0
	left, top := y, x
	for j := 0; j < SizeOfTetromino; j++ { // Search, where tetromino body starts
		if t[0][j] != '.' {
			if y < j { // If tetromino body is too big
				return false
			}
			left = y - j
			break
		}
	}

	type tetrominoNode struct{ i, j int }
	var tetrominoNodes [4]tetrominoNode
	for i := 0; i < SizeOfTetromino; i++ { // Iterating tetromino
		for j := 0; j < SizeOfTetromino; j++ {
			if t[i][j] != '.' {
				if top+i >= len(square) || left+j >= len(square) || square[top+i][left+j] != '.' { // the tetromino goes out of bounds or there is an occupied place in the square
					return false
				}
				tetrominoNodes[piecesOfTetromino] = tetrominoNode{i, j}
				piecesOfTetromino++
				if piecesOfTetromino == SizeOfTetromino { // the whole tetromino body is checked
					for n := 0; n < piecesOfTetromino; n++ {
						square[top+tetrominoNodes[n].i][left+tetrominoNodes[n].j] = t[tetrominoNodes[n].i][tetrominoNodes[n].j]
					}
					return true
				}
			}
		}
	}
	return false
}

func deleteTetrominoFromSquare(t tetromino, square [][]rune, x, y int) {
	piecesOfTetromino := 0

	left, top := y, x
	for j := 0; j < SizeOfTetromino; j++ { // Search, where tetromino body starts
		if t[0][j] != '.' {
			left = y - j
			break
		}
	}

	for i := 0; i < SizeOfTetromino; i++ { // Iterating tetromino
		for j := 0; j < SizeOfTetromino; j++ {
			if t[i][j] != '.' {
				square[top+i][left+j] = '.'
				piecesOfTetromino++
				if piecesOfTetromino == SizeOfTetromino { // the whole tetromino body is checked
					return
				}
			}
		}
	}
}
