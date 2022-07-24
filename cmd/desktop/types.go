package desktop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type agonistApp struct {
	app        fyne.App
	mainWindow fyne.Window

	mainHBox,
	mainMenuBox,
	workSpace *fyne.Container

	winElem agonistElements

	env agonistEnv
}

type agonistElements struct {
	settingsCard,
	alphabetCard,
	outdateCard,
	genSiteCard,
	aboutCard *widget.Card

	homeBtn,
	alphabetBtn,
	outdateBtn,
	genSiteBtn,
	aboutBtn *widget.Button

	settToken  *widget.Entry
	settSource *widget.SelectEntry
	settSave   *widget.Button
	settLog    *widget.Entry
	settForm   *widget.Form

	alphaTestStart *widget.Button
	alphaBar       *widget.ProgressBar
	alphaResult    *widget.Entry
	alphaBox       *fyne.Container

	outdateTestStart *widget.Button
	outdateBar       *widget.ProgressBar
	outdateResult    *widget.Entry
	outdateBox       *fyne.Container

	genSiteStart *widget.Button
	genSiteBar   *widget.ProgressBar
	genSiteLog   *widget.Entry
	genSiteBox   *fyne.Container
}

type agonistEnv struct {
	GhTiket    string `json:"Tiket"`
	SourcePath string `json:"Source"`
	IsReady    bool   `json:"Ready"`
}
