package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// afficherEcranThemes affiche la liste des themes disponibles avec la progression du joueur
func (aq *appQuiz) afficherEcranThemes() {
	salutation := creerTitre(fmt.Sprintf("Bonjour, %s !", aq.nomJoueur))

	// Barre de progression
	nbValides := len(aq.themesValides)
	nbTotal := len(aq.themes)

	barreProgression := widget.NewProgressBar()
	if nbTotal > 0 {
		barreProgression.SetValue(float64(nbValides) / float64(nbTotal))
	}

	etiquetteProgression := widget.NewLabel(fmt.Sprintf("Progression : %d / %d theme(s) valide(s)", nbValides, nbTotal))
	etiquetteProgression.Alignment = fyne.TextAlignCenter

	// Boutons pour chaque theme
	var boutonThemes []fyne.CanvasObject
	for _, t := range aq.themes {
		t := t
		if aq.themesValides[t.ID] {
			// Theme deja valide, on le desactive
			btn := widget.NewButton(t.Nom+" (valide)", nil)
			btn.Disable()
			boutonThemes = append(boutonThemes, btn)
		} else {
			btn := widget.NewButton(t.Nom, func() {
				aq.afficherEcranQuiz(t)
			})
			boutonThemes = append(boutonThemes, btn)
		}
	}

	grillleThemes := container.New(layout.NewGridLayout(2), boutonThemes...)

	boutonHistorique := widget.NewButton("Voir mon historique", func() {
		aq.afficherEcranHistorique()
	})

	boutonQuitter := widget.NewButton("Quitter", func() {
		dialog.ShowConfirm("Quitter", "Voulez-vous vraiment quitter le quiz ?", func(ok bool) {
			if ok {
				aq.fyneApp.Quit()
			}
		}, aq.fenetre)
	})
	boutonQuitter.Importance = widget.DangerImportance

	contenu := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(salutation),
		widget.NewSeparator(),
		container.NewCenter(etiquetteProgression),
		barreProgression,
		widget.NewSeparator(),
		grillleThemes,
		widget.NewSeparator(),
		boutonHistorique,
		boutonQuitter,
		layout.NewSpacer(),
	)

	aq.fenetre.SetContent(container.NewPadded(contenu))
}