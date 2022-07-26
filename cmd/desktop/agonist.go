package desktop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

// Переменная для хранения всех элементов программы
var appAgonist agonistApp

// Создание разделов приложения и их элементов.
func (a *agonistApp) createElements() {
	a.createCardSettings()
	a.createCardAlphabet()
	a.createCardOutdate()
	a.createCardGenSite()
	a.createCardAbout()

	// Создание рабочей области
	a.workSpace = container.NewGridWrap(fyne.NewSize(650, 405),
		a.winElem.settingsCard,
		a.winElem.alphabetCard,
		a.winElem.outdateCard,
		a.winElem.genSiteCard,
		a.winElem.aboutCard,
	)
}

// Создание основных оконных элементов приложения
func (a *agonistApp) createApplication() {
	a.app = app.New()

	a.mainWindow = a.app.NewWindow("[AGo]nist")
	{
		icon, _ := fyne.LoadResourceFromPath("icon.png")
		a.mainWindow.SetIcon(icon)
	}
	a.mainWindow.Resize(fyne.NewSize(700, 415))
	a.mainWindow.CenterOnScreen()
	a.mainWindow.SetFixedSize(true)

	// Создание рабочей обласит приложения всех разделов программы
	a.createElements()

	// Создание бокового меню
	a.createMenuButtons()

	// Основной горизонтальный контейнер
	a.mainHBox = container.NewHBox(a.mainMenuBox, a.workSpace)

	a.mainWindow.SetContent(a.mainHBox)
}

// Точка запуска программы
func Run() {
	// Создаем основные оконные элементы приложения
	appAgonist.createApplication()

	appAgonist.loadSettings()

	// Скрываем все разделы в рабочей области, кроме рарздела настроек
	appAgonist.winElem.settingsCard.Show()
	appAgonist.winElem.alphabetCard.Hide()
	appAgonist.winElem.outdateCard.Hide()
	appAgonist.winElem.genSiteCard.Hide()
	appAgonist.winElem.aboutCard.Hide()

	// Скрываем меню рабочих разделов
	// appAgonist.winElem.alphabetBtn.Disable()
	// appAgonist.winElem.outdateBtn.Disable()
	appAgonist.winElem.alphabetBtn.Hide()
	appAgonist.winElem.outdateBtn.Hide()
	appAgonist.winElem.genSiteBtn.Hide()

	// Запускаем оконный интерфейс
	appAgonist.mainWindow.ShowAndRun()
}
