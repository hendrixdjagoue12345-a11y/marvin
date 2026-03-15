package main
 
 
import (
	"fmt"
	"image/color"
 
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)
 
func (qa *quizApp) showHistoriqueScreen() {
	title := makeTitle(fmt.Sprintf("  Historique de %s", qa.nomJoueur))
 
	// Charger les parties (loadHistorique dans historiqueparties.go)
	hist, err := loadHistorique(qa.db, qa.joueurID)
	if err != nil {
		dialog.ShowError(err, qa.window)
		return
	}
 
	var rows []fyne.CanvasObject
 
	if len(hist) == 0 {
		rows = append(rows, widget.NewLabel("Aucune partie jouée pour l'instant."))
	} else {
		// En-tête
		rows = append(rows,
			container.New(layout.NewGridLayout(3),
				widget.NewLabelWithStyle("Thème", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabelWithStyle("Score", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
				widget.NewLabelWithStyle("Date", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}),
			),
		)
		rows = append(rows, widget.NewSeparator())
 
		// Une ligne par partie
		for _, h := range hist {
			h := h
			scoreColor := color.NRGBA{R: 200, G: 50, B: 50, A: 255}
			if h.Score >= 3 {
				scoreColor = color.NRGBA{R: 34, G: 139, B: 34, A: 255}
			}
			scoreText := canvas.NewText(fmt.Sprintf("%d/5", h.Score), scoreColor)
			scoreText.TextStyle = fyne.TextStyle{Bold: true}
			scoreText.Alignment = fyne.TextAlignCenter
 
			rows = append(rows,
				container.New(layout.NewGridLayout(3),
					widget.NewLabel(h.Theme),
					container.NewCenter(scoreText),
					widget.NewLabel(h.Date),
				),
			)
		}
	}
 
	scroll := container.NewVScroll(container.NewVBox(rows...))
	scroll.SetMinSize(fyne.NewSize(660, 320))
 
	backBtn := widget.NewButton("← Retour", func() {
		qa.showThemeScreen()
	})
	backBtn.Importance = widget.HighImportance
 
	content := container.NewVBox(
		container.NewCenter(title),
		widget.NewSeparator(),
		scroll,
		widget.NewSeparator(),
		backBtn,
	)
 
	qa.window.SetContent(container.NewPadded(content))
}
 
