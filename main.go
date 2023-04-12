package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	a := app.New()
	win := a.NewWindow("Fyne ListView Issue")
	win.SetMaster()
	win.Resize(fyne.NewSize(1024, 700))

	listSize := 300000
	listWidget := widget.NewList(
		func() int {
			return listSize
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template Object")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(strconv.Itoa(id) + " - aaaaaaa")
		},
	)

	gotoField := widget.NewEntry()
	gotoField.OnSubmitted = func(idStr string) {
		id, err := strconv.Atoi(idStr)
		if id > 0 && err == nil {
			listWidget.Select(id)
		}
		gotoField.SetText("")
	}

	win.SetContent(container.NewBorder(gotoField, nil, nil, nil, listWidget))
	win.ShowAndRun()
}
