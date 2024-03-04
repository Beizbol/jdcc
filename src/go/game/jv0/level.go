package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Level struct {
	rect  rl.Rectangle
	color rl.Color
}

func NewLevel() Level {
	return Level{
		rect:  rl.NewRectangle(0, 2*h/3, w, h/3),
		color: rl.Black,
	}
}

func (l *Level) Draw() {
	rl.DrawRectangleRec(l.rect, l.color)
}
