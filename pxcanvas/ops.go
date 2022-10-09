package pxcanvas

import (
	"fyne.io/fyne/v2"
)

func (pxCanvas *PxCanvas) Pan(previousCoords, currentCoord fyne.PointEvent) {
	xDiff := currentCoord.Position.X - previousCoords.Position.X
	yDiff := currentCoord.Position.Y - previousCoords.Position.Y

	pxCanvas.CanvasOffset.X += xDiff
	pxCanvas.CanvasOffset.Y += yDiff
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) scale(direction int) {
	switch {
	case direction > 0:
		pxCanvas.PxSize += 1
	case direction < 0:
		if pxCanvas.PxSize > 2 {
			pxCanvas.PxSize -= 1
		}
	default:
		pxCanvas.PxSize = 10
	}
}
