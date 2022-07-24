package desktop

import (
	"encoding/json"
	"os"
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
