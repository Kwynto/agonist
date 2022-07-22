package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type agonistApp struct {
	app        fyne.App
	mainWindow fyne.Window

	winElem agonistElements

	mainHBox,
	mainMenuBox,
	workSpace *fyne.Container
}

type agonistElements struct {
	blankLabel *widget.Label

	settingsCard,
	alphabetCard,
	superannuateCard,
	aboutCard *widget.Card

	homeBtn,
	alphabetBtn,
	superannuateBtn,
	aboutBtn *widget.Button

	settToken  *widget.Entry
	settSource *widget.SelectEntry
	settSave   *widget.Button
	settForm   *widget.Form
}
