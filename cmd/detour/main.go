package main

import (
	"os"
	"strings"

	"github.com/dadrian/detour/config"
	"github.com/dadrian/detour/launcher"
	"github.com/gotk3/gotk3/glib"
	"github.com/sirupsen/logrus"

	"github.com/gotk3/gotk3/gtk"
)

var c = `
browsers:
  - name: Firefox Personal
    browser: firefox
    profile: default-release
  - name: Firefox Censys
    browser: firefox
    profile: Censys
`

func main() {
	// Create Gtk Application, change appID to your application domain name reversed.
	logrus.SetLevel(logrus.DebugLevel)
	const appID = "io.dadrian.detour"
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	// Check to make sure no errors when creating Gtk Application
	if err != nil {
		logrus.Fatalf("could not create application: %s", err)
	}
	r := strings.NewReader(c)
	definitions, err := config.ParseConfig(r)
	if err != nil {
		logrus.Fatalf("error parsing config: %s", err)
	}
	if err := definitions.CheckValidity(); err != nil {
		logrus.Fatalf("error in config: %s", err)
	}

	// Application signals available
	// startup -> sets up the application when it first starts
	// activate -> shows the default first window of the application (like a new document). This corresponds to the application being launched by the desktop environment.
	// open -> opens files and shows them in a new window. This corresponds to someone trying to open a document (or documents) using the application from the file browser, or similar.
	// shutdown ->  performs shutdown tasks
	// Setup activate signal with a closure function.
	application.Connect("activate", func() {
		// Create ApplicationWindow
		appWindow, err := gtk.ApplicationWindowNew(application)
		if err != nil {
			logrus.Fatal("Could not create application window.", err)
		}
		// Set ApplicationWindow Properties
		appWindow.SetTitle("Detour")
		appWindow.SetDefaultSize(100, 100)
		box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
		appWindow.Add(box)
		if err != nil {
			logrus.Fatalf("could not create box: %s", err)
		}
		for _, browser := range definitions.Browsers {
			label := browser.Name
			button, err := gtk.ButtonNewWithLabel(label)
			if err != nil {
				logrus.Fatal("could not create button %s", label)
			}
			button.Connect("clicked", launcher.BuildLaunchCallback(&browser))
			box.PackStart(button, true, true, 0)
		}
		appWindow.ShowAll()
	})
	// Run Gtk application
	application.Run(os.Args)
}
