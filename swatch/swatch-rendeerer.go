package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type swatchRenderer struct {
	square  canvas.Rectangle
	objects []fyne.CanvasObject
	parent  *Swatch
}

func (renderer *swatchRenderer) MinSize() fyne.Size {
	return renderer.square.MinSize()
}

func (renderer *swatchRenderer) Layout(size fyne.Size) {
	renderer.objects[0].Resize(size)

}

func (renderer *swatchRenderer) Refresh() {
	renderer.Layout(fyne.NewSize(20, 20))
	renderer.square.FillColor = renderer.parent.Color
	if renderer.parent.Selected {
		renderer.square.StrokeWidth = 3
		renderer.square.StrokeColor = color.NRGBA{255, 255, 255, 255}
		renderer.objects[0] = &renderer.square
	} else {
		renderer.square.StrokeWidth = 0
		renderer.objects[0] = &renderer.square
	}
	canvas.Refresh(renderer.parent)
}

func (renderer *swatchRenderer) Objects() []fyne.CanvasObject {
	return renderer.objects
}

func (renderer *swatchRenderer) Destroy() {}
