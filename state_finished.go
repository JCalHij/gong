package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

func finished_game_update(GS *GameState, DeltaTime float32) {
	if rl.IsKeyPressed(rl.KeySpace) {
		*GS = init_game()
		GS.Update = idle_game_update
		GS.Render = idle_game_render
	}
}

func finished_game_render(GS *GameState) {
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

	// You Won / You Lost
	{
		const ResultFontSize = 70
		const WinText string = "You WON!"
		const LoseText string = "You LOST!"
		if GS.LeftScore > GS.RightScore {
			// Player won
			TextWidth := rl.MeasureText(WinText, ResultFontSize)
			rl.DrawText(WinText, (WindowWidth-TextWidth)/2, int32(float32(WindowHeight)*0.25), ResultFontSize, rl.White)
		} else {
			// Enemy won
			TextWidth := rl.MeasureText(LoseText, ResultFontSize)
			rl.DrawText(LoseText, (WindowWidth-TextWidth)/2, int32(float32(WindowHeight)*0.25), ResultFontSize, rl.White)
		}
	}
	// Instructions
	{
		const InstructionsFontSize = 30
		TextString := "Press SPACE to start a new game"
		TextWidth := rl.MeasureText(TextString, InstructionsFontSize)
		rl.DrawText(TextString, (WindowWidth-TextWidth)/2, WindowHeight-80, InstructionsFontSize, rl.White)
	}
}
