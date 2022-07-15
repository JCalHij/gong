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

type UpdateFunction func(*GameState, float32)

func idle_game_update(GS *GameState, DeltaTime float32) {

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
		var YError float32 = (GS.Ball.Y + GS.Ball.Height/2) - (GS.RightPaddle.Y - GS.RightPaddle.Height/2)
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
		// Window limits
		if BallNewPos.X+BallWidth >= float32(WindowWidth) || BallNewPos.X <= 0.0 {
			GS.BallDirection.X *= -1
		}
		if BallNewPos.Y+BallHeight >= float32(WindowHeight) || BallNewPos.Y <= 0.0 {
			GS.BallDirection.Y *= -1
		}

		// Left paddle
		var NewBallRect rl.Rectangle = rl.Rectangle{
			X:      BallNewPos.X,
			Y:      BallNewPos.Y,
			Width:  GS.Ball.Width,
			Height: GS.Ball.Height}

		if rl.CheckCollisionRecs(GS.LeftPaddle, NewBallRect) {
			GS.BallDirection.X *= -1
		}
		// Right paddle
		if rl.CheckCollisionRecs(GS.RightPaddle, NewBallRect) {
			GS.BallDirection.X *= -1
		}

		GS.Ball.X += BallSpeed * GS.BallDirection.X * DeltaTime
		GS.Ball.Y += BallSpeed * GS.BallDirection.Y * DeltaTime
	}

}

func finished_game_update(GS *GameState, DeltaTime float32) {

}

const PaddleWidth = 15
const PaddleHeight = 75
const BallWidth = 25
const BallHeight = 25
const ScoreFontSize = 85
const TextScoreSpacing = 30
const WindowWidth int32 = 1200
const WindowHeight int32 = 600
const PaddleSpeed float32 = float32(WindowHeight) * 0.3 // [px/s] Paddle speed as a percentage of the screen height
const BallSpeed float32 = float32(WindowWidth) * 0.4

var InitialLeftPaddle rl.Rectangle = rl.Rectangle{
	X:      20 + PaddleWidth,
	Y:      float32(WindowHeight-PaddleHeight) / 2.0,
	Width:  PaddleWidth,
	Height: PaddleHeight}

var InitialRightPaddle rl.Rectangle = rl.Rectangle{
	X:      float32(WindowWidth) - 20 - 2*PaddleWidth,
	Y:      float32(WindowHeight-PaddleHeight) / 2.0,
	Width:  PaddleWidth,
	Height: PaddleHeight}

var InitialBall = rl.Rectangle{
	X:      float32(WindowWidth-BallWidth) / 2.0,
	Y:      float32(WindowHeight-BallHeight) / 2.0,
	Width:  BallWidth,
	Height: BallHeight}

type GameState struct {
	LeftPaddle  rl.Rectangle
	RightPaddle rl.Rectangle
	Ball        rl.Rectangle

	BallDirection rl.Vector2

	LeftScore  int32
	RightScore int32

	Update UpdateFunction
}

func init_game() GameState {
	return GameState{
		LeftPaddle:    InitialLeftPaddle,
		RightPaddle:   InitialRightPaddle,
		Ball:          InitialBall,
		BallDirection: vec2_from_angle(rand.Float64()),
		LeftScore:     0,
		RightScore:    0,
		Update:        playing_game_update}
}

func reset_positions(GS *GameState) {
	GS.LeftPaddle = InitialLeftPaddle
	GS.RightPaddle = InitialRightPaddle
	GS.Ball = InitialBall
	GS.BallDirection = vec2_from_angle(rand.Float64())
}

func main() {
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

			// Debug performance
			rl.DrawText(fmt.Sprintf("%.03f ms", DeltaTime*1000), 10, 10, 25, rl.Red)

			rl.EndDrawing()
		}
	}

	rl.CloseWindow()
}
