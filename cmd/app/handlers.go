package app

func (a *agonistApp) homeBtn() func() {
	return func() {
		a.winElem.settingsCard.Show()
		a.winElem.alphabetCard.Hide()
		a.winElem.superannuateCard.Hide()
		a.winElem.aboutCard.Hide()
	}
}

func (a *agonistApp) alphabetBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Show()
		a.winElem.superannuateCard.Hide()
		a.winElem.aboutCard.Hide()
	}
}

func (a *agonistApp) superannuateBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Hide()
		a.winElem.superannuateCard.Show()
		a.winElem.aboutCard.Hide()
	}
}

func (a *agonistApp) aboutBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Hide()
		a.winElem.superannuateCard.Hide()
		a.winElem.aboutCard.Show()
	}
}

func SaveSettings() func() {
	return func() {
		//
	}
}
