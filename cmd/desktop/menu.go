package desktop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Creating a program menu
func (a *agonistApp) createMenuButtons() {
	icon, _ := fyne.LoadResourceFromPath("./static/img/icon-home-64.png")
	a.winElem.homeBtn = widget.NewButtonWithIcon("", icon, a.homeBtn())

	icon, _ = fyne.LoadResourceFromPath("./static/img/icon-info-64.png")
	a.winElem.alphabetBtn = widget.NewButtonWithIcon("", icon, a.alphabetBtn())

	icon, _ = fyne.LoadResourceFromPath("./static/img/icon-clock-64.png")
	a.winElem.outdateBtn = widget.NewButtonWithIcon("", icon, a.outdateBtn())

	icon, _ = fyne.LoadResourceFromPath("./static/img/icon-gensite-64.png")
	a.winElem.genSiteBtn = widget.NewButtonWithIcon("", icon, a.genSiteBtn())

	icon, _ = fyne.LoadResourceFromPath("./static/img/icon-about-64.png")
	a.winElem.aboutBtn = widget.NewButtonWithIcon("", icon, a.aboutBtn())

	// Vertical block for placing side menu buttons
	a.mainMenuBox = container.NewVBox(
		a.winElem.homeBtn,
		a.winElem.alphabetBtn,
		a.winElem.outdateBtn,
		a.winElem.genSiteBtn,
		a.winElem.aboutBtn,
	)
}
