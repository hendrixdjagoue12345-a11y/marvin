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
 
	// Un bouton par thème
	// Remplace : for _, t := range themes { fmt.Printf("%d - %s\n", t.ID, t.Nom) }
	var themeButtons []fyne.CanvasObject
	for _, t := range qa.themes {
		t := t // capture pour la closure
		if qa.themesValides[t.ID] {
			// Thème déjà validé → bouton grisé (comme les thèmes cachés en console)
			btn := widget.NewButton(fmt.Sprintf("%s", t.Nom), nil)
			btn.Disable()
			themeButtons = append(themeButtons, btn)
		} else {
			btn := widget.NewButton(fmt.Sprintf("%s", t.Nom), func() {
				qa.showQuizScreen(t)
			})
			themeButtons = append(themeButtons, btn)
		}
	}
 
	// Grille 2 colonnes
	themeGrid := container.New(layout.NewGridLayout(2), themeButtons...)
 
	// Bouton historique
	histBtn := widget.NewButton("  Voir mon historique", func() {
		qa.showHistoriqueScreen()
	})
 
	// Bouton quitter (remplace "0 - Quitter le jeu" en console)
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
 