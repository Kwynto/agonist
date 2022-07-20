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
	ch := widget.NewSelect(
		[]string{
			"Option 1",
			"Option 2",
			"Option 3",
		},
		nil,
	)
	workSpace := container.NewVBox(homeLabel, ch)

	homeIcon, _ := fyne.LoadResourceFromPath("./static/img/icon-home-64.png")
	homeBtn := widget.NewButtonWithIcon("", homeIcon, func() {
		workSpace.Objects = []fyne.CanvasObject{
			homeLabel,
			ch,
		}
	})

	alphabetIcon, _ := fyne.LoadResourceFromPath("./static/img/icon-info-64.png")
	alphabetBtn := widget.NewButtonWithIcon("", alphabetIcon, func() {
		workSpace.Objects = []fyne.CanvasObject{
			alphabetLabel,
		}
	})

	superannuateIcon, _ := fyne.LoadResourceFromPath("./static/img/icon-clock-64.png")
	superannuateBtn := widget.NewButtonWithIcon("", superannuateIcon, func() {
		workSpace.Objects = []fyne.CanvasObject{
			superannuateLabel,
		}
	})

	aboutIcon, _ := fyne.LoadResourceFromPath("./static/img/icon-about-64.png")
	aboutBtn := widget.NewButtonWithIcon("", aboutIcon, func() {})

	mainMenuBox := container.NewVBox(homeBtn, alphabetBtn, superannuateBtn, aboutBtn)

	mainHBox := container.NewHBox(mainMenuBox, workSpace)

	agonistCanvas.SetContent(mainHBox)
	agonistMainWindow.ShowAndRun()
}
