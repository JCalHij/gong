package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type OptionSelectedCallback func(*GameState)

type MenuOptionData struct {
	Name     string
	Callback OptionSelectedCallback
}

var MenuOptions [3]MenuOptionData = [3]MenuOptionData{
	{Name: "Player vs AI", Callback: on_player_vs_ai},
	{Name: "Player vs Player", Callback: on_player_vs_player},
	{Name: "Quit", Callback: on_quit},
}

func menu_update(GS *GameState, DeltaTime float32) {
	if rl.IsKeyPressed(rl.KeyS) {
		GS.SelectedOption += 1
		if GS.SelectedOption >= len(MenuOptions) {
			GS.SelectedOption = 0
		}
	}
	if rl.IsKeyPressed(rl.KeyW) {
		GS.SelectedOption -= 1
		if GS.SelectedOption < 0 {
			GS.SelectedOption = len(MenuOptions) - 1
		}
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		MenuOptions[GS.SelectedOption].Callback(GS)
	}
}

func menu_render(GS *GameState) {
	// Title
	{
		const MenuTitle string = "gong"
		const MenuTitleFontSize int32 = 115
		var TitleWidth int32 = rl.MeasureText(MenuTitle, MenuTitleFontSize)
		rl.DrawText(MenuTitle, (WindowWidth-TitleWidth)/2, 10, MenuTitleFontSize, rl.White)
	}

	// Menu options
	{
		const OptionFontSize int32 = 45
		var YPosition int32 = WindowHeight / 2
		const DeltaY int32 = OptionFontSize * 2
		for i := 0; i < len(MenuOptions); i++ {
			var TitleWidth int32 = rl.MeasureText(MenuOptions[i].Name, OptionFontSize)
			var OptionColor rl.Color = rl.Gray
			if GS.SelectedOption == i {
				OptionColor = rl.White
			}
			rl.DrawText(MenuOptions[i].Name, (WindowWidth-TitleWidth)/2, YPosition, OptionFontSize, OptionColor)
			YPosition += DeltaY
		}
	}
}

func on_player_vs_ai(GS *GameState) {

}

func on_player_vs_player(GS *GameState) {

}

func on_quit(GS *GameState) {
	GS.Running = false
}
