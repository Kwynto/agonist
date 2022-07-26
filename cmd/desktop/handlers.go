package desktop

import (
	"encoding/json"
	"os"

	"github.com/Kwynto/preserves"
)

// Хэндлер нажати на кнопку agonistApp.winElem.homeBtn в меню программы для переключения раздела
func (a *agonistApp) homeBtn() func() {
	return func() {
		a.winElem.settingsCard.Show()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Hide()
	}
}

// Хэндлер нажати на кнопку agonistApp.winElem.alphabetBtn в меню программы для переключения раздела
func (a *agonistApp) alphabetBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Show()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Hide()
	}
}

// Хэндлер нажати на кнопку agonistApp.winElem.outdateBtn в меню программы для переключения раздела
func (a *agonistApp) outdateBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Show()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Hide()
	}
}

// Хэндлер нажати на кнопку agonistApp.winElem.genSiteBtn в меню программы для переключения раздела
func (a *agonistApp) genSiteBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Show()
		a.winElem.aboutCard.Hide()
	}
}

// Хэндлер нажати на кнопку agonistApp.winElem.aboutBtn в меню программы для переключения раздела
func (a *agonistApp) aboutBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Show()
	}
}

// Хэндлер нажати на кнопку agonistApp.winElem.settSave в разделе настроек программы для сохранения настроек и предварительной загрузки данных.
func (a *agonistApp) saveSettings() func() {
	return func() {
		// Отключаем нажатую кнопку
		a.winElem.settSave.Disable()
		defer a.winElem.settSave.Enable()

		// Записываем данные из оконных элементов переменные с настройками
		a.env.GhTiket = a.winElem.settToken.Text
		a.env.SourcePath = a.winElem.settSource.Text

		// Loading a target file // FIXME: проверить
		a.outSettLog("main.zip downloading started.") // FIXME: проверить
		source := preserves.ConcatBuffer(a.env.SourcePath, "archive/refs/heads/main.zip")
		_, err := preserves.DownloadFile(source, "./data/")
		if err != nil {
			a.outSettLog("Loading main.zip failed.") // FIXME: проверить
			return
		}
		a.outSettLog("Finished loading main.zip.") // FIXME: проверить

		// Unzip
		if resUnZip := unZipReadMe(); !resUnZip {
			a.outSettLog("main.zip is incorrect.") // FIXME: проверить
		} else {
			a.outSettLog("README.md unpacked.") // FIXME: проверить
		}

		// Save enveroment to JSON // FIXME: проверить
		out, err := os.Create("./data/cfg/settings.json")
		if err != nil {
			return
		}
		encodeJSON := json.NewEncoder(out)
		err = encodeJSON.Encode(a.env)
		if err != nil {
			return
		}
		a.outSettLog("Enveroment saved.") // FIXME: проверить

		// Если всё хорошо, то завершаем подготовку и показываем кнопки меню для всех скрытых разделов
		a.env.IsReady = true
		a.winElem.alphabetBtn.Show()
		a.winElem.outdateBtn.Show()
		a.winElem.genSiteBtn.Show()
	}
}

// Хэндлер запуска тестирования алфавитного порядка
func (a *agonistApp) testAlpha() func() {
	return func() {
		a.outAlphaResult("Test.")
	}
}

// Хэндлер запуска тестирования устаревших репозиториев
func (a *agonistApp) testOutdate() func() {
	return func() {
		a.outOutdateResult("Test.")
	}
}

// Хэндлер запуска генерации web-сайта
func (a *agonistApp) generateSite() func() {
	return func() {
		a.outGenSiteLog("Test.")
	}
}
