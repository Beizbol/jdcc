package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	floor  Level
	player Player
	spike  Spike
}

const (
	w   = 800
	h   = 600
	fps = 60
)

func NewGame() Game {
	rl.InitWindow(w, h, "Go Raylib Game")
	rl.SetTargetFPS(fps)
	lvl := NewLevel()
	return Game{
		floor:  lvl,
		player: NewPlayer(w/2, lvl.rect.Y),
		spike:  NewSpike(2*w, lvl.rect.Y),
	}
}

func (g *Game) Play() {
	defer rl.CloseWindow()
	for !rl.WindowShouldClose() {
		g.update()
		g.draw()
	}
}

func (g *Game) update() {
	g.spike.Update()
	g.player.Update()
	if rl.CheckCollisionCircleRec(g.player.pos, g.player.size, g.floor.rect) {
		g.player.Land()
	}
	if rl.CheckCollisionCircles(g.player.pos, g.player.size, g.spike.pos, g.player.size) {
		g.player.Hit()
	}

}

func (g *Game) draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.DarkGray)
	g.floor.Draw()
	g.spike.Draw()
	g.player.Draw()
	rl.EndDrawing()
}
