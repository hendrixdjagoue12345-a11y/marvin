package main

 
import (
	"fmt"
 
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)
 
func (qa *quizApp) showThemeScreen() {
	greeting := makeTitle(fmt.Sprintf("Bonjour, %s !", qa.nomJoueur))
 
	// Barre de progression
	validated := len(qa.themesValides)
	total := len(qa.themes)
	progressBar := widget.NewProgressBar()
	if total > 0 {
		progressBar.SetValue(float64(validated) / float64(total))
	}
	progressLabel := widget.NewLabel(fmt.Sprintf("Progression : %d / %d thème(s) validé(s)", validated, total))
	progressLabel.Alignment = fyne.TextAlignCenter
 
	// Un bouton thème
	var themeButtons []fyne.CanvasObject
	for _, t := range qa.themes {
    t := t
    if qa.themesValides[t.ID] {
        // deja validé
        btn := widget.NewButton(t.Nom, nil)
        btn.Disable()
        themeButtons = append(themeButtons, btn)
    } else {
        btn := widget.NewButton(t.Nom, func() {
            qa.showQuizScreen(t)
        })
        themeButtons = append(themeButtons, btn)
    }
}
 
	themeGrid := container.New(layout.NewGridLayout(2), themeButtons...)
 
	histBtn := widget.NewButton("  Voir mon historique", func() {
		qa.showHistoriqueScreen()
	})
 
	quitBtn := widget.NewButton("Quitter", func() {
		dialog.ShowConfirm("Quitter", "Veux-tu vraiment quitter le quiz ?", func(ok bool) {
			if ok {
				qa.fyneApp.Quit()
			}
		}, qa.window)
	})
	quitBtn.Importance = widget.DangerImportance
 
	content := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(greeting),
		widget.NewSeparator(),
		container.NewCenter(progressLabel),
		progressBar,
		widget.NewSeparator(),
		themeGrid,
		widget.NewSeparator(),
		histBtn,
		quitBtn,
		layout.NewSpacer(),
	)
 
	qa.window.SetContent(container.NewPadded(content))
}
 