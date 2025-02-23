package main

import (
	"embed"
	"fmt"

	"github.com/RaikaSurendra/gst_application/internal/config"
	"github.com/RaikaSurendra/gst_application/internal/gst"
	"github.com/RaikaSurendra/gst_application/internal/ui"
	"github.com/wailsapp/wails/v2/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	fmt.Println("Starting GST Application...")
	config.LoadConfig()
	gst.CalculateGST()
	ui.StartUI()

	app := application.New(application.Options{
		Title:            "GST Application",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &application.RGBA{R: 255, G: 255, B: 255, A: 1},
	})

	app.Run()
}
