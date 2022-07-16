package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func menu_update(GS *GameState, DeltaTime float32) {
	if rl.IsKeyPressed(rl.KeyS) || rl.IsKeyPressed(rl.KeyDown) {
		play_sound(SFX_OptionMove)
		GS.SelectedMainMenuOption += 1
		if GS.SelectedMainMenuOption >= len(GS.MainMenuOptions) {
			GS.SelectedMainMenuOption = 0
		}
	}
	if rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyUp) {
		play_sound(SFX_OptionMove)
		GS.SelectedMainMenuOption -= 1
		if GS.SelectedMainMenuOption < 0 {
			GS.SelectedMainMenuOption = len(GS.MainMenuOptions) - 1
		}
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		play_sound(SFX_OptionSelect)
		GS.MainMenuOptions[GS.SelectedMainMenuOption].Callback(GS)
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
		for i := 0; i < len(GS.MainMenuOptions); i++ {
			var TitleWidth int32 = rl.MeasureText(GS.MainMenuOptions[i].Name, OptionFontSize)
			var OptionColor rl.Color = rl.Gray
			if GS.SelectedMainMenuOption == i {
				OptionColor = rl.White
			}
			rl.DrawText(GS.MainMenuOptions[i].Name, (WindowWidth-TitleWidth)/2, YPosition, OptionFontSize, OptionColor)
			YPosition += DeltaY
		}
	}
}

func on_player_vs_ai(GS *GameState) {
	GS.LeftInput = left_player_input
	GS.LeftPlayerHuman = true
	GS.RightInput = right_ai_input
	GS.RightPlayerHuman = false
	change_to_idle(GS)
}

func on_player_vs_player(GS *GameState) {
	GS.LeftInput = left_player_input
	GS.LeftPlayerHuman = true
	GS.RightInput = right_player_input
	GS.RightPlayerHuman = true

	change_to_idle(GS)
}

func on_quit(GS *GameState) {
	GS.Running = false
}
