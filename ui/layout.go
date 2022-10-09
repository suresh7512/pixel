package ui

import (
	"fyne.io/fyne/v2/container"
)

// It will load the all container in pixlWindow
func Setup(app *AppInit) {
	SetupMenus(app)
	swatchesContainer := BuildSwatches(app)                                                    //it will build the swatches
	colorPicker := SetupColorPicker(app)                                                       //it will build the color picker
	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker, app.PixlCanvas) //(top, bottom, left, right, center)
	app.PixlWindow.SetContent(appLayout)

}
