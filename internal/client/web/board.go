package web

import (
	"fmt"
	"log"
	"simulator/api"
)

const (
	happy    string = "(´◡`)"
	sad             = "(>_<)"
	confused        = "¯\\_(ツ)_/¯"
)

type Board struct {
	S       string // Symbol on the board
	x       int
	y       int
	Cells   [][]int
	Message string
}

func NewBoard(n int) *Board {
	// Define a 3x3 cells
	cells := make([][]int, n)
	for i := range cells {
		cells[i] = make([]int, n)
	}

	cells[0][0] = 1

	return &Board{
		S:     happy,
		x:     0,
		y:     0,
		Cells: cells}
}

func (c *Board) Print() {
	// Print the board
	fmt.Println()
	for i := range c.Cells {
		for j := range c.Cells[i] {
			fmt.Print(c.Cells[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func (c *Board) Start() {
	err := Client.Command(api.StartCommand)
	if err != nil {
		board.Err(err)
		return
	} else {
		board.OK()
	}

	c = NewBoard(len(c.Cells))
}

func (c *Board) Stop() {
	err := Client.Command(api.StopCommand)
	if err != nil {
		board.Err(err)
		return
	} else {
		board.OK()
	}

	c = NewBoard(len(c.Cells))
}

func (c *Board) Left() {
	err := Client.Command(api.LeftCommand)
	if err != nil {
		board.Err(err)
		return
	} else {
		board.OK()
	}

	if c.x-1 >= 0 {
		c.Cells[c.y][c.x] = 0
		c.x -= 1
		c.Cells[c.y][c.x] = 1
	} else {
		board.Confused()
	}
}

func (c *Board) Right() {
	err := Client.Command(api.RightCommand)
	if err != nil {
		board.Err(err)
		return
	} else {
		board.OK()
	}

	if c.x+1 < len(c.Cells) {
		c.Cells[c.y][c.x] = 0
		c.x += 1
		c.Cells[c.y][c.x] = 1
	} else {
		board.Confused()
	}
}

func (c *Board) Forward() {
	err := Client.Command(api.ForwardCommand)
	if err != nil {
		board.Err(err)
		return
	} else {
		board.OK()
	}

	if c.y-1 >= 0 {
		c.Cells[c.y][c.x] = 0
		c.y -= 1
		c.Cells[c.y][c.x] = 1
	} else {
		board.Confused()
	}
}

func (c *Board) Back() {
	err := Client.Command(api.BackCommand)
	if err != nil {
		board.Err(err)
		return
	} else {
		board.OK()
	}

	if c.y+1 < len(c.Cells) {
		c.Cells[c.y][c.x] = 0
		c.y += 1
		c.Cells[c.y][c.x] = 1
	} else {
		board.Confused()
	}
}

func (c *Board) OK() {
	c.S = happy
	c.Message = ""
}

func (c *Board) Err(err error) {
	log.Printf("%v", err)
	c.S = sad
	c.Message = err.Error()
}

func (c *Board) Confused() {
	c.S = confused
}
