package main

//go:generate go run . -o ./pkg/

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen/types"
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
		Preprocessors: []types.Preprocessor{
			types.PreprocessorFunc(func(repos gir.Repositories) {
				layershell3 := findRepository(repos, "gtk-layer-shell-0")
				layershell3.Packages = append(layershell3.Packages, gir.Package{Name: "gtk+-3.0"})

				layershell4 := findRepository(repos, "gtk4-layer-shell-0")
				layershell4.Packages = append(layershell4.Packages, gir.Package{Name: "gtk4"})
			}),
		},
	},
)

func findRepository(repos gir.Repositories, pkg string) *gir.PkgRepository {
	for i, repo := range repos {
		if repo.Pkg == pkg {
			return &repos[i]
		}
	}
	return nil
}

func main() {
	genmain.Run(data)
}
