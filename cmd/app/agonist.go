package app

// blank code for [AGo]nist

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func Run() {
	agonistApp := app.New()
	agonistMainWindow := agonistApp.NewWindow("[AGo]nist")
	agonistMainWindow.Resize(fyne.NewSize(600, 450))

	agonistIconFile, _ := fyne.LoadResourceFromPath("icon.png")
	agonistMainWindow.SetIcon(agonistIconFile)

	agonistCanvas := agonistMainWindow.Canvas()

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewRectangle(blue)

	agonistCanvas.SetContent(rect)
	agonistMainWindow.ShowAndRun()
}
