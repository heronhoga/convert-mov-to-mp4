package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/heronhoga/convert-mov-to-mp4/tools"
)

func main() {
	a := app.New()
	w := a.NewWindow("Convy by HG")

	var inputDir, outputDir string
	inputLabel := widget.NewLabel("Input folder: (none)")
	outputLabel := widget.NewLabel("Output folder: (none)")
	status := widget.NewLabel("Ready.")

	btnInput := widget.NewButton("Select Input Folder", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err == nil && list != nil {
				inputDir = list.Path()
				inputLabel.SetText("Input folder: " + inputDir)
			}
		}, w)
	})

	btnOutput := widget.NewButton("Select Output Folder", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err == nil && list != nil {
				outputDir = list.Path()
				outputLabel.SetText("Output folder: " + outputDir)
			}
		}, w)
	})

	startBtn := widget.NewButton("Start Conversion", func() {
		if inputDir == "" || outputDir == "" {
			status.SetText("Please select both folders first.")
			return
		}
		status.SetText("Converting... please wait.")
		go func() {
			tools.Convert(inputDir, outputDir)
			fyne.DoAndWait(func() {
				status.SetText("Conversion finished.")
			})
		}()
	})

	content := container.NewVBox(
		widget.NewLabel("Convy - MOV to MP4 Converter"),
		btnInput, inputLabel,
		btnOutput, outputLabel,
		startBtn,
		status,
	)

	w.Resize(fyne.NewSize(720, 540))
	w.CenterOnScreen()
	w.SetContent(content)
	w.ShowAndRun()
}
