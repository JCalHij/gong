package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

//TODO[javi]: Separate everything into multiple files (tools, game states, config/types/constants/whatever, ...)
//TODO[javi]: "Main menu" -> PvP, PvAI or Quit
//TODO[javi]: In-game pause menu -> Continue, Main Menu, Quit
//TODO[javi]: PvP / PvAI

/* Update functions */

func idle_game_update(GS *GameState, DeltaTime float32) {
	if rl.IsKeyPressed(rl.KeySpace) {
		GS.Update = playing_game_update
		GS.Render = playing_game_render
	}
}

func finished_game_update(GS *GameState, DeltaTime float32) {
	if rl.IsKeyPressed(rl.KeySpace) {
		*GS = init_game()
		GS.Update = idle_game_update
		GS.Render = idle_game_render
	}
}

/* Specific rendering functions */

func idle_game_render(GS *GameState) {
	// Instructions
	const InstructionsFontSize = 30
	TextString := "Press SPACE to begin playing"
	TextWidth := rl.MeasureText(TextString, InstructionsFontSize)
	rl.DrawText(TextString, (WindowWidth-TextWidth)/2, WindowHeight-80, InstructionsFontSize, rl.White)
}

func finished_game_render(GS *GameState) {
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

/* Game state updates */

func init_game() GameState {
	return GameState{
		LeftPaddle:    InitialLeftPaddle,
		RightPaddle:   InitialRightPaddle,
		Ball:          InitialBall,
		BallDirection: vec2_from_angle(Random.Float64()),
		LeftScore:     0,
		RightScore:    0,
		Update:        idle_game_update,
		Render:        idle_game_render}
}

func reset_positions(GS *GameState) {
	GS.LeftPaddle = InitialLeftPaddle
	GS.RightPaddle = InitialRightPaddle
	GS.Ball = InitialBall
	GS.BallDirection = vec2_from_angle(Random.Float64())
}

/* Entry Point */

func main() {
	var RandSource = rand.NewSource(time.Now().UnixNano())
	Random = rand.New(RandSource)
	rl.InitWindow(WindowWidth, WindowHeight, "gong")

	//rl.SetTargetFPS(60)
	var GS GameState = init_game()

	for !rl.WindowShouldClose() {
		var DeltaTime float32 = rl.GetFrameTime() // [s] frame time

		/* Game Logic */
		GS.Update(&GS, DeltaTime)

		/* Rendering */
		{
			rl.BeginDrawing()
			rl.ClearBackground(rl.Black)

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

			GS.Render(&GS)

			// Debug performance
			// rl.DrawText(fmt.Sprintf("%.03f ms", DeltaTime*1000), 10, 10, 25, rl.Red)

			rl.EndDrawing()
		}
	}

	rl.CloseWindow()
}
