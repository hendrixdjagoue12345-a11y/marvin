package main

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// afficherEcranAccueil affiche la page d'accueil avec les trois choix principaux
func (aq *appQuiz) afficherEcranAccueil() {
	titre := creerTitre("Quiz Go")

	sousTitre := canvas.NewText("Bienvenue ! Que souhaitez-vous faire ?", color.NRGBA{R: 80, G: 80, B: 80, A: 255})
	sousTitre.Alignment = 1 // centrer

	boutonNouvelle := widget.NewButton("Nouvelle partie", func() {
		aq.afficherEcranConnexion()
	})
	boutonNouvelle.Importance = widget.HighImportance

	boutonReprendre := widget.NewButton("Reprendre une partie", func() {
		aq.afficherEcranReprendre()
	})

	boutonQuitter := widget.NewButton("Quitter", func() {
		dialog.ShowConfirm(
			"Quitter",
			"Voulez-vous vraiment quitter le quiz ?",
			func(ok bool) {
				if ok {
					aq.fyneApp.Quit()
				}
			},
			aq.fenetre,
		)
	})
	boutonQuitter.Importance = widget.DangerImportance

	contenu := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(titre),
		container.NewCenter(sousTitre),
		widget.NewSeparator(),
		widget.NewLabel(""),
		container.NewCenter(boutonNouvelle),
		container.NewCenter(boutonReprendre),
		container.NewCenter(boutonQuitter),
		layout.NewSpacer(),
	)

	aq.fenetre.SetContent(container.NewPadded(contenu))
}