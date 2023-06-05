package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

func inputTetrominoes() ([]tetromino, error) {
	tetrominoes := []tetromino{}
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	t := tetromino{}
	lineCounter := 0
	nodeCounter := 0

	for fileScanner.Scan() {
		// Check for correct format
		line := fileScanner.Text()
		if (len(line) != SizeOfTetromino && lineCounter < SizeOfTetromino) || (len(line) != 0 && lineCounter == SizeOfTetromino) {
			return nil, errors.New("wrong size of tetromino")
		}
		char := rune(len(tetrominoes) + 'A') // the char for current tetromino
		if lineCounter == SizeOfTetromino {
			lineCounter = 0
			nodeCounter = 0
			t = tetromino{}
		} else {
			for i, rune := range line {
				switch rune {
				case '#':
					t[lineCounter][i] = char
					nodeCounter++
					if nodeCounter > SizeOfTetromino {
						return nil, errors.New("wrong numbers of nodes in a tetromino")
					}
				case '.':
					t[lineCounter][i] = '.'
				default:
					return nil, errors.New("wrong symbol in a tetromino")
				}
			}
			lineCounter++
			if lineCounter == SizeOfTetromino {
				if !isValid(t, char) { // If tetromino is not valid
					return nil, errors.New("a tetromino of a wrong shape")
				}
				t.shiftUp()
				t.shiftLeft()
				tetrominoes = append(tetrominoes, t) // Add tetromino to slice
			}
		}
	}

	return tetrominoes, nil // Return slice of tetrominoes and true
}

func isValid(t tetromino, char rune) bool {
	connections := 0 // block connections count (max 3)

	for i := 0; i < SizeOfTetromino; i++ {
		for j := 0; j < SizeOfTetromino; j++ {
			// check if there is a connection
			if t[i][j] == '.' {
				continue
			}
			switch i {
			case 0:
				if t[i+1][j] == char {
					connections++
				}
			case SizeOfTetromino - 1:
				if t[i-1][j] == char {
					connections++
				}
			default:
				if t[i-1][j] == char {
					connections++
				}
				if t[i+1][j] == char {
					connections++
				}

			}

			switch j {
			case 0:
				if t[i][j+1] == char {
					connections++
				}
			case SizeOfTetromino - 1:
				if t[i][j-1] == char {
					connections++
				}
			default:
				if t[i][j-1] == char {
					connections++
				}
				if t[i][j+1] == char {
					connections++
				}

			}
		}
	}

	if connections == 2*SizeOfTetromino || connections == 2*(SizeOfTetromino-1) {
		return true
	}
	return false

	// for i := 0; i < SizeOfTetromino; i++ {
	// 	nextRowConnection := false // next row block connection
	// 	for j := 0; j < SizeOfTetromino; j++ {
	// 		if t[i][j] == char { // If we find tetromino body
	// 			nodes++ // Add block count
	// 			if (i > 0 && t[i-1][j] == char) || (i < 3 && t[i+1][j] == char) ||
	// 				(j > 0 && t[i][j-1] == char) || (j < 3 && t[i][j+1] == char) { // If there is a block near current block
	// 				if nodes == 3 { // If there's 3 blocks already
	// 					t[i][j] = '.' // Make tetromino current coordinates empty
	// 				}
	// 				if i < 3 && t[i+1][j] == char { // If there's a block ahead in the next row, mark it
	// 					nextRowConnection = true
	// 				}
	// 				connections++
	// 			}
	// 		}
	// 	}
	// 	if nodes > 0 && nodes < SizeOfTetromino && !nextRowConnection { // If body isn't finished and there's no blocks ahead, return false
	// 		return false
	// 	}
	// 	t[i] = [SizeOfTetromino]rune{}
	// }
	// if nodes == SizeOfTetromino && connections == 3 { // If tetromino conditions met, return true
	// 	return true
	// }
	// return false
}

func (t *tetromino) shiftUp() {
	for j := 0; j < SizeOfTetromino; j++ {
		if t[0][j] != '.' { // Tetromino is already on top
			return
		}
	}
	for i := 0; i < SizeOfTetromino-1; i++ {
		t[i] = t[i+1]
	}
	// empty the last row
	for j := 0; j < SizeOfTetromino; j++ {
		t[SizeOfTetromino-1][j] = '.'
	}

	t.shiftUp()
}

func (t *tetromino) shiftLeft() {
	for i := 0; i < SizeOfTetromino; i++ {
		if t[i][0] != '.' { //  Tetromino is already on the left edge
			return
		}
	}

	for j := 0; j < SizeOfTetromino-1; j++ {
		for i := 0; i < SizeOfTetromino; i++ {
			t[i][j] = t[i][j+1]
		}
	}
	// empty the last collumnum
	for i := 0; i < SizeOfTetromino; i++ {
		t[i][SizeOfTetromino-1] = '.'
	}

	t.shiftLeft()
}
