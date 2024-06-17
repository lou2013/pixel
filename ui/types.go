package ui

import (
	"fyne.io/fyne/v2"
	"me.com/pixl/app_type"
	"me.com/pixl/swatch"
)

type AppInit struct {
	PixelWindow fyne.Window
	State       *app_type.State
	Swatches    []*swatch.Swatch
}
