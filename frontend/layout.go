package frontend

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func toggleFullscreen(mainWindow *widgets.QMainWindow) {
	if mainWindow.WindowState()&core.Qt__WindowMaximized != 0 {
		// If currently maximized, restore to normal size
		mainWindow.SetWindowState(core.Qt__WindowNoState)
	} else {
		// If not maximized, maximize the window
		mainWindow.SetWindowState(core.Qt__WindowMaximized)
	}
}

func createWindow() *widgets.QMainWindow {
	widgets.NewQApplication(len(os.Args), os.Args)

	mainWindow := widgets.NewQMainWindow(nil, 0)
	mainWindow.SetWindowFlags(core.Qt__FramelessWindowHint)
	mainWindow.Resize2(1280, 720)
	mainWindow.SetMinimumSize2(400, 300)
	mainWindow.Show()

	return mainWindow

}

func createNewTab(tabWidget *widgets.QTabWidget, tabName string) *widgets.QWidget {
	tabNumber := tabWidget.Count() + 1
	if tabName == "" {
		tabName = fmt.Sprintf("Tab %d", tabNumber)
	}

	tabContent := widgets.NewQWidget(nil, 0)

	// Create the toolbar
	toolbar, resultLabel := toolBar(tabContent)

	// Create a vertical layout for the tab content
	layout := widgets.NewQVBoxLayout2(tabContent)
	layout.AddWidget(toolbar, 0, core.Qt__AlignTop)     // Set toolbar alignment to the top
	layout.AddWidget(resultLabel, 0, core.Qt__AlignTop) // Add a stretchable empty space

	toolbar.SetStyleSheet("QToolBar { border: none; }")
	tabWidget.InsertTab(tabWidget.Count()-1, tabContent, tabName)
	if tabWidget.Count() > 1 {
		tabWidget.SetTabEnabled(tabWidget.Count()-1, false)
	}
	return tabContent
}

func toolBar(tabContent *widgets.QWidget) (*widgets.QToolBar, *widgets.QLabel) {
	toolbar := widgets.NewQToolBar2(tabContent)
	toolbar.SetToolButtonStyle(1)

	// Create toolbar buttons
	backButton := widgets.NewQPushButton2("<", toolbar)
	forwardButton := widgets.NewQPushButton2(">", toolbar)
	reloadButton := widgets.NewQPushButton2("Reload", toolbar)
	textField := widgets.NewQLineEdit(toolbar)
	goButton := widgets.NewQPushButton2("Go", toolbar)

	toolbar.AddWidget(backButton)
	toolbar.AddWidget(forwardButton)
	toolbar.AddWidget(reloadButton)
	toolbar.AddWidget(textField)
	toolbar.AddWidget(goButton)

	resultLabel := widgets.NewQLabel(nil, 0)
	resultLabel.SetTextFormat(core.Qt__PlainText)

	goButton.ConnectClicked(func(bool) {
		url := textField.Text()
		response, err := http.Get(url)
		if err != nil {
			// Handle error
			fmt.Println("Error:", err)
			return
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			// Handle error
			fmt.Println("Error:", err)
			return
		}
		resultLabel.SetText(string(body))
	})

	return toolbar, resultLabel
}
