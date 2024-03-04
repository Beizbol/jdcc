package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Spike struct {
	x0    float32
	pos   rl.Vector2
	size  float32
	speed float32
	color rl.Color
}

func NewSpike(x, y float32) Spike {
	return Spike{
		x0:    x,
		pos:   rl.NewVector2(x, y-20),
		size:  20,
		speed: 5,
		color: rl.Yellow,
	}
}

func (s *Spike) Update() {
	s.pos.X -= s.speed
	if s.pos.X < 0 {
		s.pos.X = s.x0
	}
}

func (s *Spike) Draw() {
	v1 := rl.NewVector2(s.pos.X, s.pos.Y-30)
	v2 := rl.NewVector2(s.pos.X-20, s.pos.Y+20)
	v3 := rl.NewVector2(s.pos.X+20, s.pos.Y+20)
	rl.DrawTriangle(v1, v2, v3, s.color)
}
