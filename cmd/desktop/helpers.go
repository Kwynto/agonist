package desktop

import (
	"archive/zip"
	"encoding/json"
	"io"
	"os"

	"github.com/Kwynto/preserves"
)

// Функция loadSettings() загружает сохраненные настройки из файла.
func (a *agonistApp) loadSettings() {
	// Initialization enveroment // FIXME: проверить
	a.env.GhTiket = ""
	a.env.SourcePath = ""
	a.env.IsReady = false

	// Load realy enveroment // FIXME: проверить
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

	// Set enveroment into interface // FIXME: проверить
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
	const (
		zipFile string = "./data/main.zip"
		srcFile string = "awesome-go-main/README.md"
		dstFile string = "./data/README.md"
	)

	// Распаковка содержимого архива
	zipR, err := zip.OpenReader(zipFile)
	if err != nil {
		return false
	}

	for _, file := range zipR.File {
		if file.Name == srcFile {
			r, err := file.Open()
			if err != nil {
				return false
			}
			outF, err := os.Create(dstFile)
			if err != nil {
				return false
			}
			_, err = io.Copy(outF, r)
			if err != nil {
				return false
			}
			err = r.Close()
			if err != nil {
				return false
			}
			break
		}
	}

	return true
}
