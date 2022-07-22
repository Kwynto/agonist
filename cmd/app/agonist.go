package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var appAgonist agonistApp

func (a *agonistApp) createElements() {
	// FIXME: Delete this label after develop struct
	a.winElem.blankLabel = widget.NewLabel("Blank label Blank label Blank label Blank label Blank label")

	a.createCardSettings()

	a.winElem.alphabetCard = widget.NewCard("Checking alphabetical order.", "Here you can check if the package list is in alphabetical order.", a.winElem.blankLabel)
	a.winElem.superannuateCard = widget.NewCard("Checking for outdated packages.", "A package is considered obsolete if it has not been updated for more than a year.", a.winElem.blankLabel)
	a.winElem.aboutCard = widget.NewCard("About.", "Information about the program, author and external components.", a.winElem.blankLabel)

	{
		icon, _ := fyne.LoadResourceFromPath("./static/img/icon-home-64.png")
		a.winElem.homeBtn = widget.NewButtonWithIcon("", icon, a.homeBtn())
	}
	{
		icon, _ := fyne.LoadResourceFromPath("./static/img/icon-info-64.png")
		a.winElem.alphabetBtn = widget.NewButtonWithIcon("", icon, a.alphabetBtn())
	}
	{
		icon, _ := fyne.LoadResourceFromPath("./static/img/icon-clock-64.png")
		a.winElem.superannuateBtn = widget.NewButtonWithIcon("", icon, a.superannuateBtn())
	}
	{
		icon, _ := fyne.LoadResourceFromPath("./static/img/icon-about-64.png")
		a.winElem.aboutBtn = widget.NewButtonWithIcon("", icon, a.aboutBtn())
	}
}

func (a *agonistApp) createApplication() {
	a.app = app.New()

	a.mainWindow = a.app.NewWindow("[AGo]nist")
	{
		icon, _ := fyne.LoadResourceFromPath("icon.png")
		a.mainWindow.SetIcon(icon)
	}
	a.mainWindow.Resize(fyne.NewSize(600, 450))
	a.mainWindow.CenterOnScreen()
	a.mainWindow.SetFixedSize(true)

	a.createElements()

	a.workSpace = container.NewVBox(a.winElem.settingsCard, a.winElem.alphabetCard, a.winElem.superannuateCard, a.winElem.aboutCard)
	a.mainMenuBox = container.NewVBox(a.winElem.homeBtn, a.winElem.alphabetBtn, a.winElem.superannuateBtn, a.winElem.aboutBtn)
	a.mainHBox = container.NewHBox(a.mainMenuBox, a.workSpace)

	a.mainWindow.SetContent(a.mainHBox)
}

func Run() {
	appAgonist.createApplication()

	appAgonist.winElem.settingsCard.Show()
	appAgonist.winElem.alphabetCard.Hide()
	appAgonist.winElem.superannuateCard.Hide()
	appAgonist.winElem.aboutCard.Hide()

	appAgonist.mainWindow.ShowAndRun()
}
