package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var appAgonist agonistApp

func (a *agonistApp) CreateElements() {
	// FIXME: Delete this label after develop struct
	a.winElem.blankLabel = widget.NewLabel("Blank label Blank label Blank label Blank label Blank label")

	a.winElem.settToken = widget.NewEntry()
	a.winElem.settSource = widget.NewSelectEntry([]string{
		"https://github.com/avelino/awesome-go/README.md",
		"https://github.com/Kwynto/awesome-go/README.md",
	})
	a.winElem.settSave = widget.NewButton("Save settings", SaveSettings())
	{
		item1 := widget.NewFormItem("GitHub Token:", a.winElem.settToken)
		item2 := widget.NewFormItem("Source:", a.winElem.settSource)
		item3 := widget.NewFormItem("", a.winElem.settSave)
		a.winElem.settForm = widget.NewForm(item1, item2, item3)
	}

	a.winElem.settingsCard = widget.NewCard("Settings", "You need to fill in and save the settings for the further correct operation of the program.", a.winElem.settForm)
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

func (a *agonistApp) CreateApplication() {
	a.app = app.New()

	a.mainWindow = a.app.NewWindow("[AGo]nist")
	{
		icon, _ := fyne.LoadResourceFromPath("icon.png")
		a.mainWindow.SetIcon(icon)
	}
	a.mainWindow.Resize(fyne.NewSize(600, 450))
	a.mainWindow.CenterOnScreen()
	a.mainWindow.SetFixedSize(true)

	a.CreateElements()

	a.workSpace = container.NewVBox(a.winElem.settingsCard, a.winElem.alphabetCard, a.winElem.superannuateCard, a.winElem.aboutCard)
	a.mainMenuBox = container.NewVBox(a.winElem.homeBtn, a.winElem.alphabetBtn, a.winElem.superannuateBtn, a.winElem.aboutBtn)
	a.mainHBox = container.NewHBox(a.mainMenuBox, a.workSpace)

	a.mainWindow.SetContent(a.mainHBox)
}

func Run() {
	appAgonist.CreateApplication()

	appAgonist.winElem.settingsCard.Show()
	appAgonist.winElem.alphabetCard.Hide()
	appAgonist.winElem.superannuateCard.Hide()
	appAgonist.winElem.aboutCard.Hide()

	appAgonist.mainWindow.ShowAndRun()
}
