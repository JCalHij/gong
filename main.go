package main

import (
	"fmt"
	"math"
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

func playing_game_update(GS *GameState, DeltaTime float32) {
	/* Player movement */
	{
		if rl.IsKeyDown(rl.KeyW) {
			GS.LeftPaddle.Y -= PaddleSpeed * DeltaTime
		}
		if rl.IsKeyDown(rl.KeyS) {
			GS.LeftPaddle.Y += PaddleSpeed * DeltaTime
		}
	}

	/* Enemy movement */
	{
		var YError float32 = (GS.Ball.Y + GS.Ball.Height/2) - (GS.RightPaddle.Y + GS.RightPaddle.Height/2)
		if YError > 0 {
			GS.RightPaddle.Y += PaddleSpeed * DeltaTime
		} else {
			GS.RightPaddle.Y -= PaddleSpeed * DeltaTime
		}
	}

	/* Ball movement */
	{
		var BallPos rl.Vector2 = rl.Vector2{X: GS.Ball.X, Y: GS.Ball.Y}
		var BallDeltaMovement rl.Vector2 = rl.Vector2{
			X: BallSpeed * GS.BallDirection.X * DeltaTime,
			Y: BallSpeed * GS.BallDirection.Y * DeltaTime}
		var BallNewPos = rl.Vector2Add(BallPos, BallDeltaMovement)

		// Collision checks
		// Top & Bottom Window limits
		if BallNewPos.Y+BallHeight >= float32(WindowHeight) || BallNewPos.Y <= 0.0 {
			GS.BallDirection.Y *= -1
		}
		// Left and right limits
		if BallNewPos.X+BallWidth >= float32(WindowWidth) {
			// Ball touched right side of the screen. Point for the left side.
			GS.LeftScore += 1
			reset_positions(GS)
			if GS.LeftScore >= GameWonScore {
				// Reached maximum points, you win
				GS.Update = finished_game_update
				GS.Render = finished_game_render
			} else {
				// Not yet finished, keep playing
				GS.Update = idle_game_update
				GS.Render = idle_game_render
			}
		} else if BallNewPos.X <= 0.0 {
			// Ball touched left side of the screen. Point for the right side.
			GS.RightScore += 1
			reset_positions(GS)
			if GS.RightScore >= GameWonScore {
				// Reached maximum points, you win
				GS.Update = finished_game_update
				GS.Render = finished_game_render
			} else {
				// Not yet finished, keep playing
				GS.Update = idle_game_update
				GS.Render = idle_game_render
			}
		}

		var NewBallRect rl.Rectangle = rl.Rectangle{
			X:      BallNewPos.X,
			Y:      BallNewPos.Y,
			Width:  GS.Ball.Width,
			Height: GS.Ball.Height}

		// Left paddle
		if rl.CheckCollisionRecs(GS.LeftPaddle, NewBallRect) {
			switch rect_collision_side(&GS.LeftPaddle, &GS.Ball) {
			case TopCollision, BottomCollision:
				{
					GS.BallDirection.Y *= -1
				}
			case LeftCollision:
				{
					GS.BallDirection.X *= -1
				}
			case RightCollision:
				{
					// Ball direction changes depending on where the ball was
					var MinPos float32 = GS.LeftPaddle.Y - GS.Ball.Height
					var MaxPos float32 = GS.LeftPaddle.Y + GS.LeftPaddle.Height
					var t float32 = (GS.Ball.Y - MinPos) / (MaxPos - MinPos)
					const MinAngle float32 = -math.Pi / 3 // [rad] -60 degrees
					const MaxAngle float32 = math.Pi / 3  // [rad] 60 degrees
					var NewBallAngle float32 = MinAngle*(1-t) + MaxAngle*t
					GS.BallDirection = vec2_from_angle(float64(NewBallAngle))
				}
			}
		}
		// Right paddle
		if rl.CheckCollisionRecs(GS.RightPaddle, NewBallRect) {
			switch rect_collision_side(&GS.RightPaddle, &GS.Ball) {
			case TopCollision, BottomCollision:
				{
					GS.BallDirection.Y *= -1
				}
			case RightCollision:
				{
					GS.BallDirection.X *= -1
				}
			case LeftCollision:
				{
					// Ball direction changes depending on where the ball was
					var MinPos float32 = GS.RightPaddle.Y - GS.Ball.Height
					var MaxPos float32 = GS.RightPaddle.Y + GS.RightPaddle.Height
					var t float32 = (GS.Ball.Y - MinPos) / (MaxPos - MinPos)
					const MinAngle float32 = math.Pi - math.Pi/3 // [rad] 180 - 60 degrees
					const MaxAngle float32 = math.Pi + math.Pi/3 // [rad] 180 + 60 degrees
					var NewBallAngle float32 = MinAngle*(1-t) + MaxAngle*t
					GS.BallDirection = vec2_from_angle(float64(NewBallAngle))
				}
			}
		}

		GS.Ball.X += BallSpeed * GS.BallDirection.X * DeltaTime
		GS.Ball.Y += BallSpeed * GS.BallDirection.Y * DeltaTime
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

func playing_game_render(GS *GameState) {

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

/* Tools */

func draw_rect(rect *rl.Rectangle, color rl.Color) {
	var x int32 = int32(rect.X)
	var y int32 = int32(rect.Y)
	var w int32 = int32(rect.Width)
	var h int32 = int32(rect.Height)
	rl.DrawRectangle(x, y, w, h, color)
}

func vec2_from_angle(angle float64) rl.Vector2 {
	v := rl.Vector2{
		X: float32(math.Cos(angle)),
		Y: float32(math.Sin(angle))}
	return rl.Vector2Normalize(v)
}

func rect_collision_side(Target *rl.Rectangle, Mover *rl.Rectangle) int32 {
	var IsUpperQuadrant bool = Mover.Y+Mover.Height/2 <= Target.Y+Target.Height/2
	var UpperBoundsCheck bool = Mover.Y+Mover.Height >= Target.Y
	var LowerBoundsCheck bool = Mover.Y <= Target.Y+Target.Height
	var LeftBoundsCheck bool = Mover.X+Mover.Width <= Target.X
	var RightBoundsCheck bool = Mover.X >= Target.X+Target.Width

	if UpperBoundsCheck && LowerBoundsCheck {
		if LeftBoundsCheck {
			return LeftCollision
		} else if RightBoundsCheck {
			return RightCollision
		}
	}

	if IsUpperQuadrant {
		return TopCollision
	} else {
		return BottomCollision
	}
}
