package main

import (
	"park/scenario"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Park")
	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(800, 600))
	scenario.NewScene(w)
	w.ShowAndRun()

}
