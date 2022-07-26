package desktop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// The basic structure of the application and its graphical interface.
type agonistApp struct {
	app        fyne.App
	mainWindow fyne.Window

	mainHBox,
	mainMenuBox,
	workSpace *fyne.Container

	winElem agonistElements

	env agonistEnv
}

// Type for GUI elements
type agonistElements struct {
	// Cards for storing elements of different sections of the program
	settingsCard,
	alphabetCard,
	outdateCard,
	genSiteCard,
	aboutCard *widget.Card

	// Program menu buttons
	homeBtn,
	alphabetBtn,
	outdateBtn,
	genSiteBtn,
	aboutBtn *widget.Button

	// Elements of the program settings section
	settToken  *widget.Entry
	settSource *widget.SelectEntry
	settSave   *widget.Button
	settLog    *widget.Entry
	settForm   *widget.Form

	// Alphabetical order test section elements
	alphaTestStart *widget.Button
	alphaBar       *widget.ProgressBar
	alphaResult    *widget.Entry
	alphaBox       *fyne.Container

	// Elements of the Legacy Repositories Testing Section
	outdateTestStart *widget.Button
	outdateBar       *widget.ProgressBar
	outdateResult    *widget.Entry
	outdateBox       *fyne.Container

	// Elements of the website generation section
	genSiteStart *widget.Button
	genSiteBar   *widget.ProgressBar
	genSiteLog   *widget.Entry
	genSiteBox   *fyne.Container
}

// Type to store and save settings.
type agonistEnv struct {
	GhTiket    string `json:"GhTiket"`
	SourcePath string `json:"SourcePath"`
	IsReady    bool   `json:"IsReady"`
}
