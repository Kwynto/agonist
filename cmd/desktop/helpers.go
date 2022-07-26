package desktop

import (
	"encoding/json"
	"os"

	"github.com/Kwynto/preserves"
)

// Функция loadSettings() загружает сохраненные настройки из файла.
func (a *agonistApp) loadSettings() {
	// Initialization enveroment
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

// Функция outSettLog() выводит информационное сообщение на новую строку в оконном интерфейсе в разделе настроек программы.
func (a *agonistApp) outSettLog(msg string) {
	tempText := preserves.ConcatBuffer(a.winElem.settLog.Text, msg, "\n")
	a.winElem.settLog.SetText(tempText)
}

// Функция outAlphaResult() выводит информационное сообщение на новую строку в оконном интерфейсе в разделе тестирования алфавитного порядка.
func (a *agonistApp) outAlphaResult(msg string) {
	tempText := preserves.ConcatBuffer(a.winElem.alphaResult.Text, msg, "\n")
	a.winElem.alphaResult.SetText(tempText)
}

// Функция outOutdateResult() выводит информационное сообщение на новую строку в оконном интерфейсе в разделе тестирования устаревших репозиториев
func (a *agonistApp) outOutdateResult(msg string) {
	tempText := preserves.ConcatBuffer(a.winElem.outdateResult.Text, msg, "\n")
	a.winElem.outdateResult.SetText(tempText)
}

// Функция outGenSiteLog() выводит информационное сообщение на новую строку в оконном интерфейсе в разделе генерации сайта
func (a *agonistApp) outGenSiteLog(msg string) {
	tempText := preserves.ConcatBuffer(a.winElem.genSiteLog.Text, msg, "\n")
	a.winElem.genSiteLog.SetText(tempText)
}

// Функция UnZipReadMe() распаковывает файл main.zip и извлекает из него файл README.md. Пути файлов прописаны статично в константах функции.
func unZipReadMe() bool {
	// Эта функция является адаптером для внешней функции распаковки.
	const (
		zipFile string = "./data/main.zip"
		srcFile string = "awesome-go-main/README.md"
		dstFile string = "./data/README.md"
	)

	return preserves.UnZipFile(zipFile, srcFile, dstFile)
}
