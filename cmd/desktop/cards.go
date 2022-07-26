package desktop

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Создание раздела настроек программы
func (a *agonistApp) createCardSettings() {
	a.winElem.settToken = widget.NewEntry()
	a.winElem.settToken.PlaceHolder = "Entry token for GitHub API"
	a.winElem.settSource = widget.NewSelectEntry([]string{
		"https://github.com/avelino/awesome-go/",
		"https://github.com/Kwynto/awesome-go/",
	})
	a.winElem.settSource.PlaceHolder = "Choose or entry source"
	a.winElem.settSave = widget.NewButton("Save settings and preload", a.saveSettings())
	a.winElem.settLog = widget.NewMultiLineEntry()
	a.winElem.settLog.SetMinRowsVisible(7)
	a.winElem.settLog.Disable()
	a.winElem.settLog.SetText("[AGo]nist started.\n")
	{
		item1 := widget.NewFormItem("GitHub Token:", a.winElem.settToken)
		item2 := widget.NewFormItem("Source:", a.winElem.settSource)
		item3 := widget.NewFormItem("", a.winElem.settSave)
		item4 := widget.NewFormItem("Log:", a.winElem.settLog)
		a.winElem.settForm = widget.NewForm(item1, item2, item3, item4)
	}
	a.winElem.settingsCard = widget.NewCard(
		"Settings",
		"You need to fill in and save the settings for the further correct operation of the program.",
		a.winElem.settForm,
	)
}

// Создание раздела тестирования алфавитного порядка
func (a *agonistApp) createCardAlphabet() {
	a.winElem.alphaTestStart = widget.NewButton("Start test", a.testAlpha())
	a.winElem.alphaBar = widget.NewProgressBar()
	alphaBox := container.NewGridWithColumns(2, a.winElem.alphaTestStart, a.winElem.alphaBar)
	a.winElem.alphaResult = widget.NewMultiLineEntry()
	a.winElem.alphaResult.SetMinRowsVisible(11)
	a.winElem.alphaResult.Disable()
	a.winElem.alphaBox = container.NewVBox(
		alphaBox,
		a.winElem.alphaResult,
	)
	a.winElem.alphabetCard = widget.NewCard(
		"Checking alphabetical order.",
		"Here you can check if the package list is in alphabetical order.",
		a.winElem.alphaBox,
	)
}

// Создание раздела тестирования устаревших репозиториев
func (a *agonistApp) createCardOutdate() {
	a.winElem.outdateTestStart = widget.NewButton("Start test", a.testOutdate())
	a.winElem.outdateBar = widget.NewProgressBar()
	outdateBox := container.NewGridWithColumns(2, a.winElem.outdateTestStart, a.winElem.outdateBar)
	a.winElem.outdateResult = widget.NewMultiLineEntry()
	a.winElem.outdateResult.SetMinRowsVisible(11)
	a.winElem.outdateResult.Disable()
	a.winElem.outdateBox = container.NewVBox(
		outdateBox,
		a.winElem.outdateResult,
	)
	a.winElem.outdateCard = widget.NewCard(
		"Checking for outdated packages.",
		"A package is considered obsolete if it has not been updated for more than a year.",
		a.winElem.outdateBox,
	)
}

// Создание раздела генерации web-сайта
func (a *agonistApp) createCardGenSite() {
	a.winElem.genSiteStart = widget.NewButton("Generate web site", a.generateSite())
	a.winElem.genSiteBar = widget.NewProgressBar()
	genSiteBox := container.NewGridWithColumns(2, a.winElem.genSiteStart, a.winElem.genSiteBar)
	a.winElem.genSiteLog = widget.NewMultiLineEntry()
	a.winElem.genSiteLog.SetMinRowsVisible(11)
	a.winElem.genSiteLog.Disable()
	a.winElem.genSiteBox = container.NewVBox(
		genSiteBox,
		a.winElem.genSiteLog,
	)
	a.winElem.genSiteCard = widget.NewCard(
		"Генерация вэб сайта",
		"Здесь можно сгенерировать сайт.",
		a.winElem.genSiteBox,
	)
}

// Создание раздела показа информации о программе
func (a *agonistApp) createCardAbout() {
	aboutText := widget.NewTextGrid()
	aboutText.SetText("There will be information about the author, \nthanks for the components used, \nlinks and more.\n")
	a.winElem.aboutCard = widget.NewCard(
		"About.",
		"Information about the program, author and external components.",
		aboutText,
	)
}
