package main

import (
	"image/color"
	"suresh/pixl/apptype"
	"suresh/pixl/pxcanvas"
	"suresh/pixl/swatch"
	"suresh/pixl/ui"

	"fyne.io/fyne/v2"

	//"fyne.io/fyne/app"
	"fyne.io/fyne/v2/app"
)

func main() {
	pixlApp := app.New()
	PixelWindow := pixlApp.NewWindow("pixl")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255}, //starting color of brush
		SwatchSelected: 0,
	}

	pixlCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600), //Area of pixel widget
		CanvasOffset: fyne.NewPos(0, 0),      //
		PxRows:       50,                     //actual pixel rows
		PxCols:       50,                     //actual pixel column
		PxSize:       20,                     //each pixel size
	}

	pixlCanvas := pxcanvas.NewPxCanvas(&state, pixlCanvasConfig)

	appInit := ui.AppInit{
		PixlCanvas: pixlCanvas,
		PixlWindow: PixelWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64), //total number of swatches are 64
	}

	ui.Setup(&appInit)
	appInit.PixlWindow.ShowAndRun()
}
