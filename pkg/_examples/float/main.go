package main

import (
	"log"
	"os"

	"github.com/diamondburned/gotk4-layer-shell/pkg/gtklayershell"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

func main() {
	app := gtk.NewApplication("com.github.diamondburned.gotk4-layer-shell.examples.float", 0)
	app.Connect("activate", activate)

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	if !gtklayershell.IsSupported() {
		log.Fatalln("gtk-layer-shell not supported")
	}

	appwin := gtk.NewApplicationWindow(app)
	appwin.Add(gtk.NewLabel("This window floats in the middle."))
	appwin.SetTitle("Layer Shell Example")
	appwin.SetDefaultSize(600, 300)
	appwin.Show()

	window := &appwin.Window

	gtklayershell.InitForWindow(window)
	gtklayershell.SetLayer(window, gtklayershell.LayerShellLayerTop)

	for edge := gtklayershell.Edge(0); edge < gtklayershell.LayerShellEdgeEntryNumber; edge++ {
		gtklayershell.SetAnchor(window, edge, true)
	}
}
