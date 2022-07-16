package main

import (
	"math"

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

func change_to_idle(GS *GameState) {
	GS.Update = idle_game_update
	GS.Render = idle_game_render
}

func change_to_playing(GS *GameState) {
	GS.Update = playing_game_update
	GS.Render = playing_game_render
}

func change_to_finished(GS *GameState) {
	GS.Update = finished_game_update
	GS.Render = finished_game_render
}

func change_to_menu(GS *GameState) {
	GS.SelectedOption = 0
	GS.Update = menu_update
	GS.Render = menu_render
}

func restart_game(GS *GameState) {
	reset_positions(GS)
	GS.LeftScore = 0
	GS.RightScore = 0
}

func reset_positions(GS *GameState) {
	GS.LeftPaddle = InitialLeftPaddle
	GS.RightPaddle = InitialRightPaddle
	GS.Ball = InitialBall
	GS.BallDirection = vec2_from_angle(Random.Float64())
}
