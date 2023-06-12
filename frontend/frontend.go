package frontend

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type Position struct {
	X, Y int
}

var ()

func Run() {
	mainWindow := createWindow()
	mainLayout := widgets.NewQVBoxLayout()

	tabWidget := widgets.NewQTabWidget(nil)
	mainLayout.AddWidget(tabWidget, 0, 0)
	//tabWidget.SetStyleSheet(tabStyle)
	tabWidget.DocumentMode()
	tabWidget.SetTabsClosable(true)
	tabWidget.SetMovable(true)
	tabWidget.ConnectTabCloseRequested(func(index int) {
		tabWidget.RemoveTab(index)
		if tabWidget.Count() == 0 {
			// Open a new tab
			createNewTab(tabWidget, "New Tab")
		}

	})

	exitButton := widgets.NewQPushButton2("  x  ", nil)
	fullSizeButton := widgets.NewQPushButton3(gui.QIcon_FromTheme("window-close"), "  ▭  ", nil)
	font := gui.NewQFont2("Arial", 14, 1, false)
	font2 := gui.NewQFont2("Arial", 10, 1, false)
	font3 := gui.NewQFont2("Arial", 18, 1, false)
	exitButton.SetFont(font)
	fullSizeButton.SetFont(font3)
	exitButton.SetStyleSheet(`
    QPushButton {
        background-color: transparent;
        color: black;
        border: none;
        padding: 0;
        margin-right: 4px;
    }
    QPushButton:hover {
        background-color: red;
        color: white;
    }
`)

	buttonLayout := widgets.NewQHBoxLayout()
	buttonLayout.SetContentsMargins(0, 0, 0, 0)
	buttonLayout.SetSpacing(0)

	buttonContainer := widgets.NewQWidget(nil, 0)
	buttonContainer.SetLayout(buttonLayout)
	buttonLayout.AddWidget(fullSizeButton, 0, 0)
	buttonLayout.AddWidget(exitButton, 0, 0)
	exitButton.SetFixedSize2(64, 36)
	fullSizeButton.SetFixedSize2(64, 36)
	tabWidget.SetCornerWidget(buttonContainer, core.Qt__TopRightCorner)
	tabWidget.SetStyleSheet("QTabBar::tab { height: 40px; }")
	tabWidget.IsMovable()
	exitButton.ConnectClicked(func(bool) {
		mainWindow.Close()
	})

	fullSizeButton.ConnectClicked(func(checked bool) {
		toggleFullscreen(mainWindow)
	})

	tabWidget.SetFont(font2)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(mainLayout)
	mainWindow.SetCentralWidget(widget)

	createNewTab(tabWidget, "+")            // Create initial tab
	createNewTab(tabWidget, "Original Tab") // Create "+" tab

	tabWidget.SetCurrentIndex(0) // Select the first tab initially

	widgets.QApplication_Exec()
}
