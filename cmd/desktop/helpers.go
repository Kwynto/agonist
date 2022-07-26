package desktop

import (
	"encoding/json"
	"os"

	"github.com/Kwynto/preserves"
)

// The loadSettings() function loads the saved settings from a file.
func (a *agonistApp) loadSettings() {
	// Initialization environment
	a.env.GhTiket = ""
	a.env.SourcePath = ""
	a.env.IsReady = false

	// Loading the real environment
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

	// Set environment in GUI
	a.winElem.settToken.SetText(a.env.GhTiket)
	a.winElem.settSource.SetText(a.env.SourcePath)
}

// The outSettLog() function displays an informational message on a new line in the window interface in the program settings section.
func (a *agonistApp) outSettLog(msg string) {
	tempText := preserves.ConcatBuffer(a.winElem.settLog.Text, msg, "\n")
	a.winElem.settLog.SetText(tempText)
}

// The outAlphaResult() function prints an informational message on a new line in the window interface in the alphabetical order testing section.
func (a *agonistApp) outAlphaResult(msg string) {
	tempText := preserves.ConcatBuffer(a.winElem.alphaResult.Text, msg, "\n")
	a.winElem.alphaResult.SetText(tempText)
}

// The outOutdateResult() function outputs an informational message to a new line in the window interface in the outdated repositories testing section
func (a *agonistApp) outOutdateResult(msg string) {
	tempText := preserves.ConcatBuffer(a.winElem.outdateResult.Text, msg, "\n")
	a.winElem.outdateResult.SetText(tempText)
}

// The outGenSiteLog() function displays an informational message on a new line in the window interface in the site generation section
func (a *agonistApp) outGenSiteLog(msg string) {
	tempText := preserves.ConcatBuffer(a.winElem.genSiteLog.Text, msg, "\n")
	a.winElem.genSiteLog.SetText(tempText)
}

// The UnZipReadMe() function decompresses the main.zip file and extracts the README.md file from it. File paths are written statically in function constants.
func unZipReadMe() bool {
	// This function is an adapter for an external unpack function.
	const (
		zipFile string = "./data/main.zip"
		srcFile string = "awesome-go-main/README.md"
		dstFile string = "./data/README.md"
	)

	return preserves.UnZipFile(zipFile, srcFile, dstFile)
}
