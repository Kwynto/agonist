package app

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

	blankLabel := widget.NewLabel("Blank label")

	settingsCard := widget.NewCard("Settings", "You need to fill in and save the settings for the further correct operation of the program.", blankLabel)
	alphabetCard := widget.NewCard("Checking alphabetical order.", "Here you can check if the package list is in alphabetical order.", blankLabel)
	superannuateCard := widget.NewCard("Checking for outdated packages.", "A package is considered obsolete if it has not been updated for more than a year.", blankLabel)
	aboutCard := widget.NewCard("About.", "Information about the program, author and external components.", blankLabel)

	workSpace := container.NewVBox(settingsCard, alphabetCard, superannuateCard, aboutCard)
	settingsCard.Show()
	alphabetCard.Hide()
	superannuateCard.Hide()
	aboutCard.Hide()

	homeIcon, _ := fyne.LoadResourceFromPath("./static/img/icon-home-64.png")
	homeBtn := widget.NewButtonWithIcon("", homeIcon, func() {
		settingsCard.Show()
		alphabetCard.Hide()
		superannuateCard.Hide()
		aboutCard.Hide()
	})

	alphabetIcon, _ := fyne.LoadResourceFromPath("./static/img/icon-info-64.png")
	alphabetBtn := widget.NewButtonWithIcon("", alphabetIcon, func() {
		settingsCard.Hide()
		alphabetCard.Show()
		superannuateCard.Hide()
		aboutCard.Hide()
	})

	superannuateIcon, _ := fyne.LoadResourceFromPath("./static/img/icon-clock-64.png")
	superannuateBtn := widget.NewButtonWithIcon("", superannuateIcon, func() {
		settingsCard.Hide()
		alphabetCard.Hide()
		superannuateCard.Show()
		aboutCard.Hide()
	})

	aboutIcon, _ := fyne.LoadResourceFromPath("./static/img/icon-about-64.png")
	aboutBtn := widget.NewButtonWithIcon("", aboutIcon, func() {
		settingsCard.Hide()
		alphabetCard.Hide()
		superannuateCard.Hide()
		aboutCard.Show()
	})

	mainMenuBox := container.NewVBox(homeBtn, alphabetBtn, superannuateBtn, aboutBtn)

	mainHBox := container.NewHBox(mainMenuBox, workSpace)

	agonistMainWindow.SetContent(mainHBox)
	agonistMainWindow.ShowAndRun()
}
