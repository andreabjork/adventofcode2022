package day22

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
)

func Day22(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Password: %d\n", solve(inputFile))
	} else {
		fmt.Println("Not implmenented.")
	}
}


func solve(inputFile string) int {
	board := CreateBoard(inputFile)
	for i := 0; i < len(board.instructions); i++{
		instr := board.instructions[i]
		switch instr.movement {
		case Walk:
			steps := 0
			for steps < instr.value {
				switch board.rotation {
				case Right:
					val, ok := board.surface[board.atRow][board.atCol+1]
					if !ok { // wrap around
					  col := board.atCol+2
					  for {
						  if col > board.maxCol {
							  col = 0
						  }
						  val, ok = board.surface[board.atRow][col]
						  if ok { // Found a non-empty value
						  	if val != Wall {
						  		board.atCol = col
							} // else do nothing
							break
						  }
						  col++
					  }
					} else if val != Wall {
						board.atCol++
					}
				case Down:
					val, ok := board.surface[board.atRow+1][board.atCol]
					if !ok { // wrap around
						row := board.atRow+1
						for {
							if row > board.maxRow {
								row = 0
							}
							val, ok = board.surface[row][board.atCol]
							if ok { // Found a non-empty value
								if val != Wall {
									board.atRow = row
								} // else do nothing
								break
							}
							row++
						}
					} else if val != Wall {
						board.atRow++
					}
				case Left:
					val, ok := board.surface[board.atRow][board.atCol-1]
					if !ok { // wrap around
						col := board.atCol-2
						for {
							if col < 0 {
								col = board.maxCol
							}
							val, ok = board.surface[board.atRow][col]
							if ok { // Found a non-empty value
								if val != Wall {
									board.atCol = col
								} // else do nothing
								break
							}
							col--
						}
					} else if val != Wall {
						board.atCol--
					}

				case Up:
					val, ok := board.surface[board.atRow-1][board.atCol]
					if !ok { // wrap around
						row := board.atRow-2
						for {
							if row < 0 {
								row = board.maxRow
							}
							val, ok = board.surface[row][board.atCol]
							if ok { // Found a non-empty value
								if val != Wall {
									board.atRow = row
								} // else do nothing
								break
							}
							row--
						}
					} else if val != Wall {
						board.atRow--
					}
				}

				board.surface[board.atRow][board.atCol] = board.getPath()
				steps++
			}
		case Rotate:
			if instr.value == 1 { // clockwise
				switch board.rotation {
				case Left:
					board.rotation = Up
				case Up:
					board.rotation = Right
				case Right:
					board.rotation = Down
				case Down:
					board.rotation = Left
				}
			} else { // counterclockwise
				switch board.rotation {
				case Left:
					board.rotation = Down
				case Up:
					board.rotation = Left
				case Right:
					board.rotation = Up
				case Down:
					board.rotation = Right
				}
			}
		}
	}

	board.print()
	fmt.Printf("Row,col,rot: %d,%d,%d\n", board.atRow+1, board.atCol+1, board.rotation)
	return 1000*(board.atRow+1)+4*(board.atCol+1)+int(board.rotation)
}

func (b *Board) getPath() Structure {
	switch b.rotation {
	case Right:
		return PathRight
	case Left:
		return PathLeft
	case Up:
		return PathUp
	case Down:
		return PathDown
	}
	return PathRight
}

type Rotation int64
const (
	Right 	Rotation  = 0
	Down			  = 1
	Left			  = 2
	Up 				  = 3
)
type Structure int64
const (
	Empty   Structure = 0
	Wall              = 1
	PathRight         = 2
	PathLeft          = 3
	PathDown          = 4
	PathUp            = 5
)

type Board struct {
	surface 		map[int]map[int]Structure
	instructions  	[]*Instruction
	atCol, atRow 	int
	rotation 		Rotation
	maxCol, maxRow 	int
}

func (b *Board) print() {
	for i := 0; i < b.maxRow; i++ {
		for j := 0; j < b.maxCol; j++ {
			if val, ok := b.surface[i][j]; ok {
				switch val {
				case Empty:
					fmt.Printf(".")
				case Wall:
					fmt.Printf("#")
				case PathRight:
					fmt.Printf(">")
				case PathLeft:
					fmt.Printf("<")
				case PathUp:
					fmt.Printf("^")
				case PathDown:
					fmt.Printf("v")
				}
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("---------------")
}

func CreateBoard(inputFile string) *Board {
	ls := util.LineScanner(inputFile)

	surface := map[int]map[int]Structure{}
	line, ok := util.Read(ls)

	maxRow := 0
	maxCol := 0
	rows := 0
	cols := 0
	startCol := -1
	for ok {
		surface[rows] = map[int]Structure{}
		chars := []rune(line)
		for c := 0; c < len(chars); c++ {
			switch chars[c] {
			case '.':
				if startCol == -1 {
					startCol = cols
				}
				surface[rows][cols] = Empty
			case '#':
				surface[rows][cols] = Wall
			}
			cols++
		}
		if cols > maxCol {
			maxCol = cols
		}
		cols = 0
		rows++

		line, ok = util.Read(ls)
		if line == "" {
			break
		}
	}
	if rows > maxRow {
		maxRow = rows
	}
	instructions := []*Instruction{}
	line, _ = util.Read(ls)
	instr := []rune(line)
	build := ""
	for i := 0; i < len(instr); i++ {
		if instr[i] == 'R' {
			num, err := strconv.Atoi(build)
			if err == nil {
				instructions = append(instructions, &Instruction{
					movement: Walk,
					value:  num,
				})
			}
			instructions = append(instructions, &Instruction{
				movement: Rotate,
				value:  1,
			})
			build = ""
		} else if instr[i] == 'L' {
			num, err := strconv.Atoi(build)
			if err == nil {
				instructions = append(instructions, &Instruction{
					movement: Walk,
					value:  num,
				})
			}
			instructions = append(instructions, &Instruction{
				movement: Rotate,
				value:  -1,
			})
			build = ""
		} else {
			build += string(instr[i])
		}
	}

	num, err := strconv.Atoi(build)
	if err == nil {
		instructions = append(instructions, &Instruction{
			movement: Walk,
			value:  num,
		})
	}
	fmt.Printf("cols %d, rows %d\n", maxCol, maxRow)
	return &Board{
		surface:      surface,
		instructions: instructions,
		atCol: startCol,
		atRow: 0,
		rotation: Right,
		maxCol: maxCol,
		maxRow: maxRow,
	}
}

type Movement int64
const (
	Walk 	 Movement = 0
	Rotate 			  = 1
)

type Instruction struct {
	movement 	Movement
	value 		int
}
