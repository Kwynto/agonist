package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type agonistApp struct {
	app        fyne.App
	mainWindow fyne.Window

	blankLabel *widget.Label

	settingsCard,
	alphabetCard,
	superannuateCard,
	aboutCard *widget.Card

	homeBtn,
	alphabetBtn,
	superannuateBtn,
	aboutBtn *widget.Button

	mainHBox,
	mainMenuBox,
	workSpace *fyne.Container
}

var appAgonist agonistApp

func (a *agonistApp) NewApplication() {
	a.app = app.New()

	a.mainWindow = a.app.NewWindow("[AGo]nist")
	{
		icon, _ := fyne.LoadResourceFromPath("icon.png")
		a.mainWindow.SetIcon(icon)
		a.mainWindow.CenterOnScreen()
	}
	a.mainWindow.Resize(fyne.NewSize(600, 450))

	a.blankLabel = widget.NewLabel("Blank label Blank label Blank label Blank label Blank label")

	a.settingsCard = widget.NewCard("Settings", "You need to fill in and save the settings for the further correct operation of the program.", a.blankLabel)
	a.alphabetCard = widget.NewCard("Checking alphabetical order.", "Here you can check if the package list is in alphabetical order.", a.blankLabel)
	a.superannuateCard = widget.NewCard("Checking for outdated packages.", "A package is considered obsolete if it has not been updated for more than a year.", a.blankLabel)
	a.aboutCard = widget.NewCard("About.", "Information about the program, author and external components.", a.blankLabel)

	a.workSpace = container.NewVBox(a.settingsCard, a.alphabetCard, a.superannuateCard, a.aboutCard)

	{
		icon, _ := fyne.LoadResourceFromPath("./static/img/icon-home-64.png")
		a.homeBtn = widget.NewButtonWithIcon("", icon, func() {
			a.settingsCard.Show()
			a.alphabetCard.Hide()
			a.superannuateCard.Hide()
			a.aboutCard.Hide()
		})
	}
	{
		icon, _ := fyne.LoadResourceFromPath("./static/img/icon-info-64.png")
		a.alphabetBtn = widget.NewButtonWithIcon("", icon, func() {
			a.settingsCard.Hide()
			a.alphabetCard.Show()
			a.superannuateCard.Hide()
			a.aboutCard.Hide()
		})
	}
	{
		icon, _ := fyne.LoadResourceFromPath("./static/img/icon-clock-64.png")
		a.superannuateBtn = widget.NewButtonWithIcon("", icon, func() {
			a.settingsCard.Hide()
			a.alphabetCard.Hide()
			a.superannuateCard.Show()
			a.aboutCard.Hide()
		})
	}
	{
		icon, _ := fyne.LoadResourceFromPath("./static/img/icon-about-64.png")
		a.aboutBtn = widget.NewButtonWithIcon("", icon, func() {
			a.settingsCard.Hide()
			a.alphabetCard.Hide()
			a.superannuateCard.Hide()
			a.aboutCard.Show()
		})
	}
	a.mainMenuBox = container.NewVBox(a.homeBtn, a.alphabetBtn, a.superannuateBtn, a.aboutBtn)

	a.mainHBox = container.NewHBox(a.mainMenuBox, a.workSpace)

	a.mainWindow.SetContent(a.mainHBox)
}

func Run() {
	appAgonist.NewApplication()

	appAgonist.settingsCard.Show()
	appAgonist.alphabetCard.Hide()
	appAgonist.superannuateCard.Hide()
	appAgonist.aboutCard.Hide()

	appAgonist.mainWindow.ShowAndRun()
}
