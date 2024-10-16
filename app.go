package main

import (
	utils "changeme/backend"
	"changeme/backend/csgsi"
	steamutils "changeme/backend/steamUtils"
	"context"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"syscall"

	"github.com/lxn/win"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx  context.Context
	mode string
}

// NewApp creates a new App application struct
func NewApp(mode string) *App {
	return &App{
		mode: mode,
	}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx

	if a.mode == "hud" {
		hwnd := win.FindWindow(nil, syscall.StringToUTF16Ptr("HUD"))
		win.SetWindowLong(hwnd, win.GWL_EXSTYLE, win.GetWindowLong(hwnd, win.GWL_EXSTYLE)|win.WS_EX_LAYERED)
	}
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here

	if a.mode != "hud" {
		return
	}
	runtime.LogDebug(ctx, "Stating GSI Services")

	gsi := csgsi.New(ctx)

	go func() {
		runtime.LogDebug(ctx, "gsi-raw-state thread started")
		runtime.EventsEmit(ctx, "gsi-ready", true)

		for state := range gsi.RawChannel {
			runtime.EventsEmit(ctx, "gsi-raw-state", state)
		}
	}()

	go func(gs *csgsi.Game) {
		// iterates over all channels and prints the value
		v := reflect.ValueOf(gs.Events)
		t := v.Type()

		for i := 0; i < v.NumField(); i++ {
			fieldName := t.Field(i).Name
			field := v.Field(i)

			if field.Kind() == reflect.Chan {
				runtime.LogDebug(ctx, "gsi-"+fieldName+" thread started")
				go func(name string, ch reflect.Value) {
					for {
						value, ok := ch.Recv()
						if !ok {
							break
						}
						var serializedValue interface{}
						if value.Kind() == reflect.Struct {
							serializedValue = value.Interface()
						} else {
							serializedValue = fmt.Sprintf("%v", value)
						}

						runtime.EventsEmit(ctx, name, serializedValue)
					}
				}(fieldName, field)
			}
		}
	}(gsi)

	gsi.Listen(ctx, ":3000")
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetMode() string {
	runtime.LogDebug(a.ctx, "Found mode "+a.mode)

	return a.mode
}

func (a *App) OpenSelf(mode string) error {
	path, _ := os.Executable()

	runtime.LogDebug(a.ctx, "Opening self with mode: "+mode)
	var cmd string
	var args []string
	cmd = path
	args = []string{"--mode", mode}

	return exec.Command(cmd, args...).Start()
}

func (a *App) GetSteamAvatar(steamId string) string {
	// example xml link: https://steamcommunity.com/profiles/76561197995773793/?xml=1
	url := "https://steamcommunity.com/profiles/" + steamId + "/?xml=1"

	// now we need to fetch the xml
	xmlBytes, err := utils.GetXML(url)
	if err != nil {
		return ""
	}

	var result steamutils.ProfileXML
	xml.Unmarshal(xmlBytes, &result)

	return result.AvatarFull
}
