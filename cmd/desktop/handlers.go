package desktop

import (
	"encoding/json"
	"os"

	"github.com/Kwynto/preserves"
)

// Handler click on the agonistApp.winElem.homeBtn button in the program menu to switch the section
func (a *agonistApp) homeBtn() func() {
	return func() {
		a.winElem.settingsCard.Show()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Hide()
	}
}

// Handler click on the button agonistApp.winElem.alphabetBtn in the program menu to switch the section
func (a *agonistApp) alphabetBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Show()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Hide()
	}
}

// Handler click on the agonistApp.winElem.outdateBtn button in the program menu to switch the section
func (a *agonistApp) outdateBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Show()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Hide()
	}
}

// Handler click on the agonistApp.winElem.genSiteBtn button in the program menu to switch the section
func (a *agonistApp) genSiteBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Show()
		a.winElem.aboutCard.Hide()
	}
}

// Handler click on the button agonistApp.winElem.aboutBtn in the program menu to switch the section
func (a *agonistApp) aboutBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Show()
	}
}

// Handler click on the agonistApp.winElem.settSave button in the program settings section to save settings and preload data.
func (a *agonistApp) saveSettings() func() {
	return func() {
		// Disable the pressed button
		a.winElem.settSave.Disable()
		defer a.winElem.settSave.Enable()

		// We write data from window elements variables with settings
		a.env.GhTiket = a.winElem.settToken.Text
		a.env.SourcePath = a.winElem.settSource.Text

		// Loading a target file
		a.outSettLog("main.zip downloading started.")
		source := preserves.ConcatBuffer(a.env.SourcePath, "archive/refs/heads/main.zip")
		_, err := preserves.DownloadFile(source, "./data/")
		if err != nil {
			a.outSettLog("Loading main.zip failed.")
			return
		}
		a.outSettLog("Finished loading main.zip.")

		// Unzip
		if resUnZip := unZipReadMe(); !resUnZip {
			a.outSettLog("main.zip is incorrect.")
		} else {
			a.outSettLog("README.md unpacked.")
		}

		// Save environment to JSON
		out, err := os.Create("./data/cfg/settings.json")
		if err != nil {
			return
		}
		encodeJSON := json.NewEncoder(out)
		err = encodeJSON.Encode(a.env)
		if err != nil {
			return
		}
		a.outSettLog("Enveroment saved.")

		// If everything is fine, then we complete the preparation and show the menu buttons for all hidden sections
		a.env.IsReady = true
		a.winElem.alphabetBtn.Show()
		a.winElem.outdateBtn.Show()
		a.winElem.genSiteBtn.Show()
	}
}

// Alphabetical Test Trigger Handler
func (a *agonistApp) testAlpha() func() {
	return func() {
		a.outAlphaResult("Test.")
	}
}

// Run handler for testing obsolete repositories
func (a *agonistApp) testOutdate() func() {
	return func() {
		a.outOutdateResult("Test.")
	}
}

// Web site generation startup handler
func (a *agonistApp) generateSite() func() {
	return func() {
		a.outGenSiteLog("Test.")
	}
}
