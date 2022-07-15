package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/gen2brain/raylib-go/raylib"
)

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

const PaddleWidth = 30
const PaddleHeight = 75
const BallWidth = 25
const BallHeight = 25
const ScoreFontSize = 85
const TextScoreSpacing = 30

func main() {
	var WindowWidth int32 = 1200
	var WindowHeight int32 = 600

	var PaddleSpeed float32 = float32(WindowHeight) * 0.3 // [px/s] Paddle speed as a percentage of the screen height
	var BallSpeed float32 = float32(WindowWidth) * 0.2

	rl.InitWindow(WindowWidth, WindowHeight, "gong")

	//rl.SetTargetFPS(60)

	LeftPaddle := rl.Rectangle{
		X:      20 + PaddleWidth,
		Y:      float32(WindowHeight-PaddleHeight) / 2.0,
		Width:  PaddleWidth,
		Height: PaddleHeight}
	var LeftScore = 0

	RightPaddle := rl.Rectangle{
		X:      float32(WindowWidth) - 20 - 2*PaddleWidth,
		Y:      float32(WindowHeight-PaddleHeight) / 2.0,
		Width:  PaddleWidth,
		Height: PaddleHeight}
	var RightScore = 0

	Ball := rl.Rectangle{
		X:      float32(WindowWidth-BallWidth) / 2.0,
		Y:      float32(WindowHeight-BallHeight) / 2.0,
		Width:  BallWidth,
		Height: BallHeight}

	var BallDirection rl.Vector2 = vec2_from_angle(rand.Float64())

	for !rl.WindowShouldClose() {
		var DeltaTime float32 = rl.GetFrameTime() // [s] frame time

		/* Game Logic */

		RightScoreText := fmt.Sprintf("%d", RightScore)
		var RightTextWidth = rl.MeasureText(RightScoreText, ScoreFontSize)

		LeftScoreText := fmt.Sprintf("%d", LeftScore)

		/* Player movement */
		{
			if rl.IsKeyDown(rl.KeyW) {
				LeftPaddle.Y -= PaddleSpeed * DeltaTime
			}
			if rl.IsKeyDown(rl.KeyS) {
				LeftPaddle.Y += PaddleSpeed * DeltaTime
			}
		}

		/* Enemy movement */
		{
			//TODO
		}

		/* Ball movement */
		{
			var BallPos rl.Vector2 = rl.Vector2{X: Ball.X, Y: Ball.Y}
			var BallDeltaMovement rl.Vector2 = rl.Vector2{
				X: BallSpeed * BallDirection.X * DeltaTime,
				Y: BallSpeed * BallDirection.Y * DeltaTime}
			var BallNewPos = rl.Vector2Add(BallPos, BallDeltaMovement)

			// Collision checks
			// Window limits
			if BallNewPos.X+BallWidth >= float32(WindowWidth) || BallNewPos.X <= 0.0 {
				BallDirection.X *= -1
			}
			if BallNewPos.Y+BallHeight >= float32(WindowHeight) || BallNewPos.Y <= 0.0 {
				BallDirection.Y *= -1
			}

			Ball.X += BallSpeed * BallDirection.X * DeltaTime
			Ball.Y += BallSpeed * BallDirection.Y * DeltaTime
		}

		/* Rendering */
		{
			rl.BeginDrawing()
			rl.ClearBackground(rl.Black)

			// Paddles
			draw_rect(&LeftPaddle, rl.White)
			draw_rect(&RightPaddle, rl.White)
			// Ball
			draw_rect(&Ball, rl.White)

			// Score
			rl.DrawText(LeftScoreText, WindowWidth/2.0-TextScoreSpacing-RightTextWidth, 10, ScoreFontSize, rl.White)
			rl.DrawText(RightScoreText, WindowWidth/2.0+TextScoreSpacing, 10, ScoreFontSize, rl.White)

			// Debug performance
			rl.DrawText(fmt.Sprintf("%.03f ms", DeltaTime*1000), 10, 10, 25, rl.Red)

			rl.EndDrawing()
		}
	}

	rl.CloseWindow()
}
