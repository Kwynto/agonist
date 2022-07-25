package desktop

import (
	"encoding/json"
	"os"

	"github.com/Kwynto/preserves"
)

func (a *agonistApp) loadSettings() {
	// Initialization enveroment
	a.env.GhTiket = ""
	a.env.SourcePath = ""
	a.env.IsReady = false

	// Load realy enveroment
	var inEnv agonistEnv
	in, err := os.Open("./data/cfg/settings.json")
	if err != nil {
		return
	}
	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(&inEnv)
	if err != nil {
		return
	}
	in.Close()
	a.env = inEnv

	// Set enveroment into interface
	a.winElem.settToken.SetText(a.env.GhTiket)
	a.winElem.settSource.SetText(a.env.SourcePath)
}

func (a *agonistApp) outSettLog(msg string) {
	tempText := preserves.ConcatBuffer(a.winElem.settLog.Text, msg, "\n")
	a.winElem.settLog.SetText(tempText)
}

func (a *agonistApp) outAlphaResult(msg string) {
	tempText := preserves.ConcatBuffer(a.winElem.alphaResult.Text, msg, "\n")
	a.winElem.alphaResult.SetText(tempText)
}

func (a *agonistApp) outOutdateResult(msg string) {
	tempText := preserves.ConcatBuffer(a.winElem.outdateResult.Text, msg, "\n")
	a.winElem.outdateResult.SetText(tempText)
}

func (a *agonistApp) outGenSiteLog(msg string) {
	tempText := preserves.ConcatBuffer(a.winElem.genSiteLog.Text, msg, "\n")
	a.winElem.genSiteLog.SetText(tempText)
}
