package main

//go:generate go run . -o ./pkg/

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
)

const layerShellModule = "github.com/diamondburned/gotk4-layer-shell/pkg"

var data = genmain.Overlay(
	gendata.Main,
	genmain.Data{
		Module: layerShellModule,
		Packages: []genmain.Package{
			{Name: "gtk-layer-shell-0", Namespaces: nil},
			{Name: "gtk4-layer-shell-0", Namespaces: nil},
		},
		PkgGenerated: []string{
			"gtklayershell",
			"gtk4layershell",
		},
		PkgExceptions: []string{
			"go.mod",
			"go.sum",
			"LICENSE",
			"_examples",
		},
	},
)

func main() {
	genmain.Run(data)
}
