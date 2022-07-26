package desktop

// Variable to store all elements of the program
var appAgonist agonistApp

// Program start point
func Run() {
	// Creating the main window elements of the application
	appAgonist.createApplication()

	appAgonist.loadSettings()

	// Hide all sections in the workspace, except for the settings section
	appAgonist.winElem.settingsCard.Show()
	appAgonist.winElem.alphabetCard.Hide()
	appAgonist.winElem.outdateCard.Hide()
	appAgonist.winElem.genSiteCard.Hide()
	appAgonist.winElem.aboutCard.Hide()

	// Hiding the menu of working sections
	// appAgonist.winElem.alphabetBtn.Disable()
	// appAgonist.winElem.outdateBtn.Disable()
	appAgonist.winElem.alphabetBtn.Hide()
	appAgonist.winElem.outdateBtn.Hide()
	appAgonist.winElem.genSiteBtn.Hide()

	// Launching the window interface
	appAgonist.mainWindow.ShowAndRun()
}
