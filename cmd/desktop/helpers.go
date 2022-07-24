package desktop

import (
	"encoding/json"
	"os"
)

func (a *agonistApp) loadSettings() {
	// Initialization enveroment
	a.env.ghTiket = ""
	a.env.sourcePath = ""
	a.env.isReady = false

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
	a.winElem.settToken.SetText(a.env.ghTiket)
	a.winElem.settSource.SetText(a.env.sourcePath)
}
