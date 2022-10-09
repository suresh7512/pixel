package ui

import (
	//"suresh/pixel/apptype"
	//"suresh/pixel/swatch"
	"suresh/pixl/apptype"
	"suresh/pixl/pxcanvas"
	"suresh/pixl/swatch"

	//"fyne.io/fyne"
	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
