package fyneKit

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var NewApp func() fyne.App = app.New

var NewVBox func(objects ...fyne.CanvasObject) *fyne.Container = container.NewVBox

var NewSize func(w float32, h float32) fyne.Size = fyne.NewSize

var NewLabel func(text string) *widget.Label = widget.NewLabel

var NewButton func(label string, tapped func()) *widget.Button = widget.NewButton
