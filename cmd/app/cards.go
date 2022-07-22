package app

import (
	"fyne.io/fyne/v2/widget"
)

func (a *agonistApp) createCardSettings() {
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
}
