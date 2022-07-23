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
	developLabel *widget.Label

	settingsCard,
	alphabetCard,
	outdateCard,
	aboutCard *widget.Card

	homeBtn,
	alphabetBtn,
	outdateBtn,
	aboutBtn *widget.Button

	settToken  *widget.Entry
	settSource *widget.SelectEntry
	settSave   *widget.Button
	settForm   *widget.Form

	alphaTestStart *widget.Button
	alphaBar       *widget.ProgressBar
	alphaResult    *widget.Entry
	alphaBox       *fyne.Container

	outdateTestStart *widget.Button
	outdateBar       *widget.ProgressBar
	outdateResult    *widget.Entry
	outdateBox       *fyne.Container
}
