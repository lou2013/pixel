package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"me.com/pixl/app_type"
	"me.com/pixl/pxcanvas"
	"me.com/pixl/swatch"
	"me.com/pixl/ui"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")

	state := app_type.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	pixlCanvasConfig := app_type.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       10,
		PxCols:       10,
		PxSize:       30,
	}

	pixlCanvas := pxcanvas.NewPxCanvas(&state, pixlCanvasConfig)

	appInit := ui.AppInit{
		PixlCanvas:  pixlCanvas,
		PixelWindow: pixlWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixelWindow.ShowAndRun()
}
