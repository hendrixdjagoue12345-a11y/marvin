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

// afficherEcranHistorique affiche l'historique des parties du joueur
func (aq *appQuiz) afficherEcranHistorique() {
	titre := creerTitre(fmt.Sprintf("Historique de %s", aq.nomJoueur))

	historique, err := chargerHistorique(aq.db, aq.joueurID)
	if err != nil {
		dialog.ShowError(err, aq.fenetre)
		return
	}

	var lignes []fyne.CanvasObject

	if len(historique) == 0 {
		lignes = append(lignes, widget.NewLabel("Aucune partie jouee pour l'instant."))
	} else {
		// En-tete du tableau
		lignes = append(lignes,
			container.New(layout.NewGridLayout(3),
				widget.NewLabelWithStyle("Theme", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabelWithStyle("Score", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
				widget.NewLabelWithStyle("Date", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}),
			),
		)
		lignes = append(lignes, widget.NewSeparator())

		// Une ligne par partie
		for _, h := range historique {
			h := h
			couleurScore := color.NRGBA{R: 200, G: 50, B: 50, A: 255}
			if h.Score >= 3 {
				couleurScore = color.NRGBA{R: 34, G: 139, B: 34, A: 255}
			}
			texteScore := canvas.NewText(fmt.Sprintf("%d/5", h.Score), couleurScore)
			texteScore.TextStyle = fyne.TextStyle{Bold: true}
			texteScore.Alignment = fyne.TextAlignCenter

			lignes = append(lignes,
				container.New(layout.NewGridLayout(3),
					widget.NewLabel(h.Theme),
					container.NewCenter(texteScore),
					widget.NewLabel(h.Date),
				),
			)
		}
	}

	defilement := container.NewVScroll(container.NewVBox(lignes...))
	defilement.SetMinSize(fyne.NewSize(660, 320))

	boutonRetour := widget.NewButton("Retour", func() {
		aq.afficherEcranThemes()
	})
	boutonRetour.Importance = widget.HighImportance

	boutonPDF := widget.NewButton("Telecharger PDF", func() {
		dialog.NewConfirm(
			"Telechargement PDF",
			"Voulez-vous telecharger un PDF de l'historique ?",
			func(ok bool) {
				if ok {
					err := genererPDF(historique, aq.nomJoueur)
					if err != nil {
						dialog.ShowError(err, aq.fenetre)
						return
					}
					dialog.ShowInformation("Succes", "PDF genere avec succes.", aq.fenetre)
				}
			},
			aq.fenetre,
		).Show()
	})
	boutonPDF.Importance = widget.SuccessImportance

	boutons := container.NewHBox(boutonRetour, layout.NewSpacer(), boutonPDF)

	contenu := container.NewVBox(
		container.NewCenter(titre),
		widget.NewSeparator(),
		defilement,
		widget.NewSeparator(),
		boutons,
	)

	aq.fenetre.SetContent(container.NewPadded(contenu))
}