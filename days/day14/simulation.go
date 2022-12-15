package day14

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"time"
)

func VisualizeA(inputFile string) {
	c := MakeCave(inputFile, false)
	vis := &Visualization{
		width: c.maxX-c.minX,
		height: c.maxY,

		score:  0,
		isOver: false,
	}

	p := &Path{[]Step{}, 500, 0}
	OOB, _ := c.fall(p)

	// Do the termbox visualization:
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	//go listenToKeyboard(keyboardEventsChan)
	if err := vis.render(c); err != nil {
		panic(err)
	}

	i := 1
	for !OOB {
		if err := vis.render(c.trace(p)); err != nil {
			panic(err)
		}
		time.Sleep(vis.moveInterval())
		p, OOB = c.next(p)
		if !OOB {
			i++
		}
	}
}

func VisualizeB(inputFile string) {
	c := MakeCave(inputFile, true)
	vis := &Visualization{
		width: c.maxX-c.minX,
		height: c.maxY-350,
		score:  0,
		isOver: false,
	}

	p := &Path{[]Step{}, 500, 0}
	_, _ = c.fall(p)

	// Do the termbox visualization:
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	//go listenToKeyboard(keyboardEventsChan)
	if err := vis.render(c.trace(p)); err != nil {
		panic(err)
	}

	i := 1
	for !(p.endX == 500 && p.endY == 0) {
		if err := vis.render(c.trace(p)); err != nil {
			panic(err)
		}
		time.Sleep(vis.moveInterval())
		p, _ = c.next(p)
		i++
	}
}

// Game type
type Visualization struct {
	width 	int
	height 	int
	score  int
	isOver bool
}

func (vis *Visualization) moveInterval() time.Duration {
	ms := 3
	return time.Duration(ms) * time.Millisecond
}

const (
	defaultColor = termbox.ColorDefault
	rockColor 	 = termbox.ColorDarkGray
	bgColor      = termbox.ColorDefault
	snakeColor   = termbox.ColorYellow
)

func (vis *Visualization) render(c *Cave) error {
	termbox.Clear(defaultColor, defaultColor)

	var (
		w, h   = termbox.Size()
		midY   = h / 2
		left   = (w - vis.width) / 2
		//right  = (w + vis.width) / 2
		//top    = midY - (vis.height / 2)
		bottom = midY + (vis.height / 2) + 1
	)

	renderCave(left, bottom, c)

	return termbox.Flush()
}

func renderCave(left, bottom int, c *Cave) {
	for y := 0; y <= c.maxY; y++ {
		for x := c.minX; x <= c.maxX; x++ {
			switch _, s := c.get(x,y); s {
			case Rock:
				termbox.SetCell(left+x-c.minX, bottom+y, '█', rockColor, snakeColor)
			case Sand:
				termbox.SetCell(left+x-c.minX, bottom+y, '*', snakeColor, bgColor)
			case Trail:
				termbox.SetCell(left+x-c.minX, bottom+y, '~', snakeColor, bgColor)
			}
		}
	}
}

func renderTitle(left, top int) {
	tbprint(left, top-1, defaultColor, defaultColor, "Snake Game")
}

func renderScore(left, bottom, s int) {
	score := fmt.Sprintf("Score: %v", s)
	tbprint(left, bottom+1, defaultColor, defaultColor, score)
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func (vis *Visualization) end() {
	vis.isOver = true
}
