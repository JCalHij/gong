package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

func idle_game_update(GS *GameState, DeltaTime float32) {
	if rl.IsKeyPressed(rl.KeySpace) {
		change_to_playing(GS)
	}
}

func idle_game_render(GS *GameState) {
	// Paddles
	draw_rect(&GS.LeftPaddle, rl.White)
	draw_rect(&GS.RightPaddle, rl.White)
	// Ball
	draw_rect(&GS.Ball, rl.White)

	// Score
	{
		RightScoreText := fmt.Sprintf("%d", GS.RightScore)
		var RightTextWidth = rl.MeasureText(RightScoreText, ScoreFontSize)

		LeftScoreText := fmt.Sprintf("%d", GS.LeftScore)
		rl.DrawText(LeftScoreText, WindowWidth/2.0-TextScoreSpacing-RightTextWidth, 10, ScoreFontSize, rl.White)
		rl.DrawText(RightScoreText, WindowWidth/2.0+TextScoreSpacing, 10, ScoreFontSize, rl.White)
	}

	// Instructions
	const InstructionsFontSize = 30
	TextString := "Press SPACE to begin playing"
	TextWidth := rl.MeasureText(TextString, InstructionsFontSize)
	rl.DrawText(TextString, (WindowWidth-TextWidth)/2, WindowHeight-80, InstructionsFontSize, rl.White)

	if GS.LeftPlayerHuman {
		// "Move up and down with keys W and S"
		TextString = "Move up and down"
		rl.DrawText(TextString, 10, WindowHeight/2-PaddleHeight*2, InstructionsFontSize, rl.Gray)
		TextString = "with keys W and S"
		rl.DrawText(TextString, 10, WindowHeight/2-PaddleHeight*2+InstructionsFontSize, InstructionsFontSize, rl.Gray)
	}

	if GS.RightPlayerHuman {
		// "Move up and down with arrows up and down"
		TextString = "Move up and down"
		TextWidth = rl.MeasureText(TextString, InstructionsFontSize)
		rl.DrawText(TextString, WindowWidth-TextWidth-20, WindowHeight/2-PaddleHeight*2, InstructionsFontSize, rl.Gray)
		TextString = "with arrows up and down"
		TextWidth = rl.MeasureText(TextString, InstructionsFontSize)
		rl.DrawText(TextString, WindowWidth-TextWidth-20, WindowHeight/2-PaddleHeight*2+InstructionsFontSize, InstructionsFontSize, rl.Gray)
	}
}
