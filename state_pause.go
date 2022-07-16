package main

import rl "github.com/gen2brain/raylib-go/raylib"

func pause_update(GS *GameState, DeltaTime float32) {
	if rl.IsKeyPressed(rl.KeyEscape) {
		change_to_playing(GS)
		return
	}
}

func pause_render(GS *GameState) {
	draw_paddles_ball_and_score(GS)
}
