package main

import (
	"math/rand"

	"github.com/gen2brain/raylib-go/raylib"
)

type UpdateFunction func(*GameState, float32)
type RenderFunction func(*GameState)

type InputFunction func(*GameState, float32)

type OptionSelectedCallback func(*GameState)

type MenuOptionData struct {
	Name     string
	Callback OptionSelectedCallback
}

type GameState struct {
	// App state

	Update UpdateFunction
	Render RenderFunction

	Running bool

	// Main game stuff

	LeftPaddle  rl.Rectangle
	RightPaddle rl.Rectangle
	Ball        rl.Rectangle

	LeftInput        InputFunction
	LeftPlayerHuman  bool
	RightInput       InputFunction
	RightPlayerHuman bool

	BallDirection rl.Vector2

	LeftScore  int32
	RightScore int32

	SelectedPauseMenuOption int
	PauseMenuOptions        [2]MenuOptionData

	// Menu stuff

	SelectedMainMenuOption int
	MainMenuOptions        [3]MenuOptionData
}

const PaddleWidth = 15
const PaddleHeight = 75
const BallWidth = 25
const BallHeight = 25
const ScoreFontSize = 85
const TextScoreSpacing = 30
const WindowWidth int32 = 1200
const WindowHeight int32 = 600
const PaddleSpeed float32 = float32(WindowHeight) * 0.35 // [px/s] Paddle speed as a percentage of the screen height
const BallSpeed float32 = float32(WindowWidth) * 0.45
const GameWonScore int32 = 5

const (
	None            = iota
	TopCollision    = iota
	LeftCollision   = iota
	BottomCollision = iota
	RightCollision  = iota
)

//Sounds
const (
	SFX_PaddleHit    = iota
	SFX_WallHit      = iota
	SFX_Goal         = iota
	SFX_OptionMove   = iota
	SFX_OptionSelect = iota
	SFX_MAX          = iota
)

var SFX_Resources [SFX_MAX]string = [SFX_MAX]string{
	"assets/ball_paddle_hit.wav",
	"assets/wall_hit.wav",
	"assets/goal.wav",
	"assets/option_move.wav",
	"assets/option_selected.wav",
}

var SFX_Sounds [SFX_MAX]rl.Sound

var Random *rand.Rand

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
