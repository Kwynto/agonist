package app

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (a *agonistApp) createCardSettings() {
	a.winElem.settToken = widget.NewEntry()
	a.winElem.settSource = widget.NewSelectEntry([]string{
		"https://github.com/avelino/awesome-go/README.md",
		"https://github.com/Kwynto/awesome-go/README.md",
	})
	a.winElem.settSave = widget.NewButton("Save settings", a.saveSettings())
	{
		item1 := widget.NewFormItem("GitHub Token:", a.winElem.settToken)
		item2 := widget.NewFormItem("Source:", a.winElem.settSource)
		item3 := widget.NewFormItem("", a.winElem.settSave)
		a.winElem.settForm = widget.NewForm(item1, item2, item3)
	}
	a.winElem.settingsCard = widget.NewCard("Settings", "You need to fill in and save the settings for the further correct operation of the program.", a.winElem.settForm)
}

func (a *agonistApp) createCardAlphabet() {
	a.winElem.alphaTestStart = widget.NewButton("Start test", a.testAlpha())
	a.winElem.alphaBar = widget.NewProgressBar()
	alphaBox := container.NewGridWithColumns(2, a.winElem.alphaTestStart, a.winElem.alphaBar)
	a.winElem.alphaResult = widget.NewMultiLineEntry()
	a.winElem.alphaResult.SetMinRowsVisible(11)
	a.winElem.alphaBox = container.NewVBox(
		alphaBox,
		a.winElem.alphaResult,
	)
	a.winElem.alphabetCard = widget.NewCard("Checking alphabetical order.", "Here you can check if the package list is in alphabetical order.", a.winElem.alphaBox)
}
