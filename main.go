package main

import "fmt"
import "github.com/gen2brain/raylib-go/raylib"

func draw_rect(rect *rl.Rectangle, color rl.Color) {
	var x int32 = int32(rect.X)
	var y int32 = int32(rect.Y)
	var w int32 = int32(rect.Width)
	var h int32 = int32(rect.Height)
	rl.DrawRectangle(x, y, w, h, color)
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

	var DeltaTime float32 = 0.0

	for !rl.WindowShouldClose() {
		DeltaTime = rl.GetFrameTime()

		/* Game Logic */

		RightScoreText := fmt.Sprintf("%d", RightScore)
		var RightTextWidth = rl.MeasureText(RightScoreText, ScoreFontSize)

		LeftScoreText := fmt.Sprintf("%d", LeftScore)

		/* Rendering */

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
		rl.DrawText(fmt.Sprintf("%.03f ms", DeltaTime*1000), 10, 10, 25, rl.White)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
