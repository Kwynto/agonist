package desktop

import (
	"encoding/json"
	"os"
)

func (a *agonistApp) homeBtn() func() {
	return func() {
		a.winElem.settingsCard.Show()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Hide()
	}
}

func (a *agonistApp) alphabetBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Show()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Hide()
	}
}

func (a *agonistApp) outdateBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Show()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Hide()
	}
}

func (a *agonistApp) genSiteBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Show()
		a.winElem.aboutCard.Hide()
	}
}

func (a *agonistApp) aboutBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Show()
	}
}

func (a *agonistApp) saveSettings() func() {
	return func() {
		// var outEnv agonistEnv

		a.env.GhTiket = a.winElem.settToken.Text
		a.env.SourcePath = a.winElem.settSource.Text
		// outEnv = a.env

		out, err := os.Create("./data/cfg/settings.json")
		if err != nil {
			return
		}
		encodeJSON := json.NewEncoder(out)
		err = encodeJSON.Encode(a.env)
		if err != nil {
			return
		}
	}
}

func (a *agonistApp) testAlpha() func() {
	return func() {
		//
	}
}

func (a *agonistApp) testOutdate() func() {
	return func() {
		//
	}
}

func (a *agonistApp) generateSite() func() {
	return func() {
		//
	}
}
