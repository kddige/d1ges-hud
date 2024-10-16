package main

import (
	"embed"
	"flag"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

var mode = flag.String("mode", "hud", "The running mode of the application, accepted: normal, hud")

func main() {

	flag.Parse()
	// Create an instance of the app structure
	app := NewApp(*mode)

	// Create application with options

	var opts *options.App
	if *mode == "hud" {
		opts = &options.App{
			Title:            "HUD",
			Width:            1024,
			Height:           768,
			Assets:           assets,
			BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
			Fullscreen:       true,
			AlwaysOnTop:      false,
			DisableResize:    true,
			Frameless:        true,
			OnStartup:        app.startup,
			OnDomReady:       app.domReady,
			Bind: []interface{}{
				app,
			},
		}
	} else {
		opts = &options.App{
			Title:            "wails-events",
			Width:            1024,
			Height:           768,
			Assets:           assets,
			BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
			OnStartup:        app.startup,
			OnDomReady:       app.domReady,
			Bind: []interface{}{
				app,
			},
		}
	}

	err := wails.Run(opts)

	if err != nil {
		println("Error:", err.Error())
	}
}
