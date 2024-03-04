package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	pos   rl.Vector2
	vel   rl.Vector2
	size  float32
	speed float32
	color rl.Color
	state State
	timer int
}

type State int

func NewPlayer(x, y float32) Player {
	return Player{
		pos:   rl.NewVector2(x, y-20),
		size:  20,
		speed: 5,
		color: rl.Blue,
	}
}

// Player states
const ( // go enums lol
	_     = iota // 0
	stand        // 1
	hurt         // 2
)

func (p *Player) Hit() {
	p.state = hurt
}

func (p *Player) Land() {
	p.state = stand
}

const grav = 0.1

func (p *Player) Update() {
	p.timer++
	p.vel.X = 0
	if rl.IsKeyDown(rl.KeyW) {
		p.vel.Y = -p.speed
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.vel.Y = p.speed
	}
	if rl.IsKeyDown(rl.KeyA) {
		p.vel.X = -p.speed
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.vel.X = p.speed
	}
	p.vel.Y += grav
	p.pos.X += p.vel.X
	p.pos.Y += p.vel.Y
	switch p.state {
	case stand:
		if p.timer > 60 {
			p.state = 0
		}
		p.color = rl.Red
	case hurt:
		if p.timer > 60 {
			p.state = 0
		}
		p.color = rl.Red
	default:
		p.color = rl.Blue
		p.timer = 0
	}
}

func (p *Player) Draw() {
	rl.DrawCircleV(p.pos, p.size, p.color)
}
