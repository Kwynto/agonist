package desktop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

// Creation of application sections and their elements.
func (a *agonistApp) createElements() {
	a.createCardSettings()
	a.createCardAlphabet()
	a.createCardOutdate()
	a.createCardGenSite()
	a.createCardAbout()

	// Create a workspace
	a.workSpace = container.NewGridWrap(fyne.NewSize(650, 405),
		a.winElem.settingsCard,
		a.winElem.alphabetCard,
		a.winElem.outdateCard,
		a.winElem.genSiteCard,
		a.winElem.aboutCard,
	)
}

// Creating the main window elements of the application
func (a *agonistApp) createApplication() {
	a.app = app.New()

	a.mainWindow = a.app.NewWindow("[AGo]nist")
	{
		icon, _ := fyne.LoadResourceFromPath("icon.png")
		a.mainWindow.SetIcon(icon)
	}
	a.mainWindow.Resize(fyne.NewSize(700, 415))
	a.mainWindow.CenterOnScreen()
	a.mainWindow.SetFixedSize(true)

	// Creation of the working area of the application of all sections of the program
	a.createElements()

	// Creating a side menu
	a.createMenuButtons()

	// Main horizontal container
	a.mainHBox = container.NewHBox(a.mainMenuBox, a.workSpace)

	a.mainWindow.SetContent(a.mainHBox)
}
