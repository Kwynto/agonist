package app

// blank code for [AGo]nist

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	agonistApp := app.New()
	agonistMainWindow := agonistApp.NewWindow("[AGo]nist")
	agonistMainWindow.Resize(fyne.NewSize(600, 450))

	agonistIconFile, _ := fyne.LoadResourceFromPath("icon.png")
	agonistMainWindow.SetIcon(agonistIconFile)
	agonistMainWindow.CenterOnScreen()

	agonistCanvas := agonistMainWindow.Canvas()

	homeLabel := widget.NewLabel("Home")
	alphabetLabel := widget.NewLabel("Alphabet order")
	superannuateLabel := widget.NewLabel("Superannuate")
	clarificationLabel := widget.NewLabel("Clarification")
	ch := widget.NewSelect(
		[]string{
			"Option 1",
			"Option 2",
			"Option 3",
		},
		nil,
	)
	workSpace := container.NewVBox(homeLabel, ch)

	homeBtn := widget.NewButton("Home", func() {
		workSpace.Objects = []fyne.CanvasObject{
			homeLabel,
			ch,
		}
	})
	alphabet := widget.NewButton("Alphabet order", func() {
		workSpace.Objects = []fyne.CanvasObject{
			alphabetLabel,
		}
	})
	superannuate := widget.NewButton("Superannuate", func() {
		workSpace.Objects = []fyne.CanvasObject{
			superannuateLabel,
		}
	})
	clarification := widget.NewButton("Clarification", func() {
		workSpace.Objects = []fyne.CanvasObject{
			clarificationLabel,
		}
	})
	mainMenuBox := container.NewVBox(homeBtn, alphabet, superannuate, clarification)

	mainHBox := container.NewHBox(mainMenuBox, workSpace)

	agonistCanvas.SetContent(mainHBox)
	agonistMainWindow.ShowAndRun()
}
