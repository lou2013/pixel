package brush

import (
	"fyne.io/fyne/v2/driver/desktop"
	"me.com/pixl/app_type"
)

const (
	Pixel = iota
)

func TryPaintPixel(appState *app_type.State, canvas app_type.Brushable, ev *desktop.MouseEvent) bool {
	x, y := canvas.MouseToCanvasXY(ev)
	if x != nil && y != nil && ev.Button == desktop.MouseButtonPrimary {
		canvas.SetColor(appState.BrushColor, *x, *y)
		return true
	}
	return false
}

func TryBrush(appState *app_type.State, canvas app_type.Brushable, ev *desktop.MouseEvent) bool {
	switch {
	case appState.BrushType == Pixel:
		return TryPaintPixel(appState, canvas, ev)
	default:
		return false
	}
}
