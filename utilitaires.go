package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// creerTitre cree un titre stylise utilise sur tous les ecrans
func creerTitre(texte string) *canvas.Text {
	t := canvas.NewText(texte, color.NRGBA{R: 0, G: 120, B: 215, A: 255})
	t.TextSize = 26
	t.TextStyle = fyne.TextStyle{Bold: true}
	t.Alignment = fyne.TextAlignCenter
	return t
}