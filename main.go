package main

import (
	"github.com/murlokswarm/app"
	_ "github.com/murlokswarm/mac"
)

var (
	win app.Contexter
)

func main() {
	app.OnLaunch = func() {
		appMenu := &AppMainMenu{}    // Creates the AppMainMenu component.
		menu, ok := app.MenuBar()
		if !ok {
			return
		}
		menu.Mount(appMenu) // Mounts the AppMainMenu component into the application menu bar.

		appMenuDock := &AppMainMenu{} // Creates another AppMainMenu.
		dock, ok := app.Dock()
		if !ok {
			return
		}
		dock.Mount(appMenuDock)

		win = newMainWindow() // Create the main window.
	}

	app.Run()
}

func newMainWindow() app.Contexter {
	// Creates a window context.
	win := app.NewWindow(app.Window{
		Title:          "Hello World",
		Width:          1280,
		Height:         720,
		TitlebarHidden: true,
		OnClose: func() bool {
			win = nil
			return true
		},
	})

	hello := &Hello{} // Creates a Hello component.
	win.Mount(hello)  // Mounts the Hello component into the window context.
	return win
}
