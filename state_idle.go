package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func idle_game_update(GS *GameState, DeltaTime float32) {
	if rl.IsKeyPressed(rl.KeySpace) {
		GS.Update = playing_game_update
		GS.Render = playing_game_render
	}
}

func idle_game_render(GS *GameState) {
	// Instructions
	const InstructionsFontSize = 30
	TextString := "Press SPACE to begin playing"
	TextWidth := rl.MeasureText(TextString, InstructionsFontSize)
	rl.DrawText(TextString, (WindowWidth-TextWidth)/2, WindowHeight-80, InstructionsFontSize, rl.White)
}
