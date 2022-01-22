package main

import (
	"bytes"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const boardSize = 16
const numBombs = 40
const bombBoardPositionValue = 99

type BoardCell struct {
	value   int
	visible bool
	color   string
}

type Board struct {
	positions [boardSize][boardSize]BoardCell
}

type BoardPosition struct {
	y int
	x int
}

func randomizeBombs() []BoardPosition {
	var positions [boardSize * boardSize]BoardPosition
	counter := 0

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			positions[counter] = BoardPosition{y: y, x: x}
			counter++
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(positions), func(i, j int) { positions[i], positions[j] = positions[j], positions[i] })

	return positions[0:numBombs]
}

func (board Board) print() {
	var byteBuffer bytes.Buffer

	byteBuffer.WriteString("\n")

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			var position = board.positions[i][j]

			if position.value == bombBoardPositionValue {
				byteBuffer.WriteString(strconv.Itoa(position.value))
			} else {
				byteBuffer.WriteString(" " + strconv.Itoa(position.value))
			}

			if position.visible {
				byteBuffer.WriteString(" ")
			} else {
				byteBuffer.WriteString("H")
			}
			byteBuffer.WriteString("|")
		}
		byteBuffer.WriteString("\n")
	}

	log.Println(byteBuffer.String())
}

func findSurroundingPositions(y int, x int) []BoardPosition {
	var neighbors []BoardPosition

	yNotTop := y != 0
	yNotBottom := y != (boardSize - 1)

	xNotLeft := x != 0
	xNotRight := x != (boardSize - 1)

	// Over
	if yNotTop && xNotLeft {
		// Diagonal top left
		neighbors = append(neighbors, BoardPosition{y: y - 1, x: x - 1})
	}
	if yNotTop {
		// Directly over
		neighbors = append(neighbors, BoardPosition{y: y - 1, x: x})
	}
	if yNotTop && xNotRight {
		// Diagonal top right
		neighbors = append(neighbors, BoardPosition{y: y - 1, x: x + 1})
	}

	// Sides
	if xNotLeft {
		// Left side
		neighbors = append(neighbors, BoardPosition{y: y, x: x - 1})
	}
	if xNotRight {
		// Right side
		neighbors = append(neighbors, BoardPosition{y: y, x: x + 1})
	}

	// Under
	if yNotBottom && xNotLeft {
		// Diagonal under left
		neighbors = append(neighbors, BoardPosition{y: y + 1, x: x - 1})
	}
	if yNotBottom {
		// Directly under
		neighbors = append(neighbors, BoardPosition{y: y + 1, x: x})
	}
	if yNotBottom && xNotRight {
		// Diagonal under right
		neighbors = append(neighbors, BoardPosition{y: y + 1, x: x + 1})
	}

	return neighbors
}

func calculateBoardValue(y int, x int, positions [boardSize][boardSize]BoardCell) int {
	surroundingPositions := findSurroundingPositions(y, x)

	neighborsBombs := 0
	for _, position := range surroundingPositions {
		if positions[position.y][position.x].value == bombBoardPositionValue {
			neighborsBombs++
		}
	}

	return neighborsBombs
}

func queueContainsPosition(y int, x int, queue []BoardPosition) bool {
	for _, position := range queue {
		if position.y == y && position.x == x {
			return true
		}
	}

	return false
}

func registerClick(board *Board, y int, x int) {
	positions := &board.positions
	positions[0][0].visible = true

	positions[y][x].visible = true

	if positions[y][x].value != 0 {
		return
	}

	log.Println("Do stuff")
	queue := make([]BoardPosition, 0)
	queue = append(queue, findSurroundingPositions(y, x)...)

	log.Println(len(queue))

	for len(queue) > 0 {
		currentPosition := queue[0]
		currentElement := &positions[currentPosition.y][currentPosition.x]
		queue = queue[1:]

		if currentElement.value == bombBoardPositionValue {
			continue
		}

		currentElement.visible = true

		if currentElement.value == 0 {
			newPositions := findSurroundingPositions(currentPosition.y, currentPosition.x)

			for _, newPosition := range newPositions {
				if !positions[newPosition.y][newPosition.x].visible && !queueContainsPosition(newPosition.y, newPosition.x, queue) {
					queue = append(queue, newPosition)
				}
			}
		}
	}
}

func createBoard() {
	var positions [boardSize][boardSize]BoardCell
	bombPositions := randomizeBombs()

	// Add bomb positions
	for _, bombPosition := range bombPositions {
		positions[bombPosition.y][bombPosition.x] = BoardCell{value: bombBoardPositionValue}
	}

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			var position = positions[y][x]
			if position.value != bombBoardPositionValue {
				var value = calculateBoardValue(y, x, positions)
				positions[y][x] = BoardCell{value: value}
			}
		}
	}

	var board = Board{positions: positions}
	board.print()

	log.Println("")
	log.Println("")
	log.Println("")

	registerClick(&board, 5, 5)
	board.print()
}
