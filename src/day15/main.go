package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	y int
	x int
}

type Game struct {
	robot Pos
	grid  [][]rune
	moves []Pos
}

func moveToDirection(move rune) Pos {
	switch move {
	case 'v':
		return Pos{1, 0}
	case '^':
		return Pos{-1, 0}
	case '>':
		return Pos{0, 1}
	case '<':
		return Pos{0, -1}
	default:
		return Pos{0, 0}
	}
}

func parseMoves(game Game, moves string) Game {
	for _, move := range moves {
		dir := moveToDirection(move)
		if dir.y == 0 && dir.x == 0 {
			continue
		}
		game.moves = append(game.moves, dir)
	}
	return game
}

func parse() (game Game) {
	data1, _ := os.ReadFile("data/day15_part1.txt")
	data2, _ := os.ReadFile("data/day15_part2.txt")
	lines := strings.Split(strings.TrimSpace(string(data1)), "\n")
	moves := string(data2)
	for y, line := range lines {
		var row []rune
		for x, c := range line {
			if c == '@' {
				game.robot.y = y
				game.robot.x = x
				row = append(row, '.')
				continue
			}
			row = append(row, c)
		}
		game.grid = append(game.grid, row)
	}
	return parseMoves(game, moves)
}

func parse2() (game Game) {
	data1, _ := os.ReadFile("data/day15_part1.txt")
	data2, _ := os.ReadFile("data/day15_part2.txt")
	lines := strings.Split(strings.TrimSpace(string(data1)), "\n")
	moves := string(data2)

	for y, line := range lines {
		var row []rune
		for x, c := range line {
			if c == 'O' {
				row = append(row, '[', ']')
				continue
			} else if c == '#' {
				row = append(row, '#', '#')
				continue
			}
			if c == '@' {
				game.robot.y = y
				game.robot.x = x * 2
			}
			row = append(row, '.', '.')
		}
		game.grid = append(game.grid, row)
	}
	return parseMoves(game, moves)
}

func solve(game Game) (score int) {
	for _, move := range game.moves {
		if !canMove(game.robot, game, move) {
			continue
		}
		doMove(game.robot, game, move)
		game.robot.x += move.x
		game.robot.y += move.y
	}
	for y, row := range game.grid {
		for x, c := range row {
			if c != '[' && c != 'O' {
				continue
			}
			score += y*100 + x
		}
	}

	return score
}

func canMove(pos Pos, game Game, dir Pos) bool {
	pos.y += dir.y
	pos.x += dir.x
	tag := game.grid[pos.y][pos.x]

	switch tag {
	case 'O':
		return canMove(pos, game, dir)
	case ']':
		return canMove(pos, game, dir) && (dir.x == -1 || dir.x == 1 || canMove(Pos{pos.y, pos.x - 1}, game, dir))
	case '[':
		return canMove(pos, game, dir) && (dir.x == -1 || dir.x == 1 || canMove(Pos{pos.y, pos.x + 1}, game, dir))
	case '.':
		return true
	default:
		return false
	}
}

func doMove(pos Pos, game Game, dir Pos) {
	if !canMove(pos, game, dir) {
		return
	}
	mem := pos
	pos.y += dir.y
	pos.x += dir.x
	tag := game.grid[pos.y][pos.x]
	updown := dir.y == -1 || dir.y == 1
	if tag == ']' || tag == '[' || tag == 'O' {
		doMove(pos, game, dir)
	}
	if tag == ']' && updown {
		doMove(Pos{pos.y, pos.x - 1}, game, dir)
	}
	if tag == '[' && updown {
		doMove(Pos{pos.y, pos.x + 1}, game, dir)
	}
	game.grid[pos.y][pos.x] = game.grid[mem.y][mem.x]
	game.grid[mem.y][mem.x] = '.'
}

func main() {
	answer1 := solve(parse())
	answer2 := solve(parse2())
	fmt.Println("ANSWER 1: ", answer1)
	fmt.Println("ANSWER 2: ", answer2)
}
