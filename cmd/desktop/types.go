package desktop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Основная структура приложения и его графического интерфейса.
type agonistApp struct {
	app        fyne.App
	mainWindow fyne.Window

	mainHBox,
	mainMenuBox,
	workSpace *fyne.Container

	winElem agonistElements

	env agonistEnv
}

// Тип для элементов графического интерфейса
type agonistElements struct {
	// Карты для хранения элементов разных разделов программы
	settingsCard,
	alphabetCard,
	outdateCard,
	genSiteCard,
	aboutCard *widget.Card

	// Кнопки меню программы
	homeBtn,
	alphabetBtn,
	outdateBtn,
	genSiteBtn,
	aboutBtn *widget.Button

	// Элементы раздела настроек программы
	settToken  *widget.Entry
	settSource *widget.SelectEntry
	settSave   *widget.Button
	settLog    *widget.Entry
	settForm   *widget.Form

	// Элементы раздела тестирования алфавитного порядка
	alphaTestStart *widget.Button
	alphaBar       *widget.ProgressBar
	alphaResult    *widget.Entry
	alphaBox       *fyne.Container

	// Элементы раздела тестирования устаревших репозиториев
	outdateTestStart *widget.Button
	outdateBar       *widget.ProgressBar
	outdateResult    *widget.Entry
	outdateBox       *fyne.Container

	// Элементы раздела генерации web-сайта
	genSiteStart *widget.Button
	genSiteBar   *widget.ProgressBar
	genSiteLog   *widget.Entry
	genSiteBox   *fyne.Container
}

// Тип для хранения и сохранения настроек.
type agonistEnv struct {
	GhTiket    string `json:"GhTiket"`
	SourcePath string `json:"SourcePath"`
	IsReady    bool   `json:"IsReady"`
}
