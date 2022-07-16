package main

import (
	"fmt"
	"math"

	"github.com/gen2brain/raylib-go/raylib"
)

func playing_game_update(GS *GameState, DeltaTime float32) {
	/* Input from playerAI */

	GS.LeftInput(GS, DeltaTime)
	GS.RightInput(GS, DeltaTime)

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
				change_to_finished(GS)
			} else {
				// Not yet finished, keep playing
				change_to_idle(GS)
			}
		} else if BallNewPos.X <= 0.0 {
			// Ball touched left side of the screen. Point for the right side.
			GS.RightScore += 1
			reset_positions(GS)
			if GS.RightScore >= GameWonScore {
				// Reached maximum points, you win
				change_to_finished(GS)
			} else {
				// Not yet finished, keep playing
				change_to_idle(GS)
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

func playing_game_render(GS *GameState) {
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
}

func left_player_input(GS *GameState, DeltaTime float32) {
	if rl.IsKeyDown(rl.KeyW) {
		GS.LeftPaddle.Y -= PaddleSpeed * DeltaTime
	}
	if rl.IsKeyDown(rl.KeyS) {
		GS.LeftPaddle.Y += PaddleSpeed * DeltaTime
	}
}

func right_player_input(GS *GameState, DeltaTime float32) {
	if rl.IsKeyDown(rl.KeyUp) {
		GS.RightPaddle.Y -= PaddleSpeed * DeltaTime
	}
	if rl.IsKeyDown(rl.KeyDown) {
		GS.RightPaddle.Y += PaddleSpeed * DeltaTime
	}
}

func right_ai_input(GS *GameState, DeltaTime float32) {
	var YError float32 = (GS.Ball.Y + GS.Ball.Height/2) - (GS.RightPaddle.Y + GS.RightPaddle.Height/2)
	if YError > 0 {
		GS.RightPaddle.Y += PaddleSpeed * DeltaTime
	} else {
		GS.RightPaddle.Y -= PaddleSpeed * DeltaTime
	}
}
