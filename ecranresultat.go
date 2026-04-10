package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// afficherEcranResultat affiche le score obtenu apres un quiz
func (aq *appQuiz) afficherEcranResultat(t theme, score, total int) {
	_ = enregistrerPartie(aq.db, aq.joueurID, t.ID, score)

	valide := score >= 3
	if valide {
		aq.themesValides[t.ID] = true
	}

	var texteTitre string
	if valide {
		texteTitre = "Theme valide !"
	} else {
		texteTitre = "Score insuffisant"
	}
	titre := creerTitre(texteTitre)

	couleurScore := color.NRGBA{R: 200, G: 50, B: 50, A: 255}
	if valide {
		couleurScore = color.NRGBA{R: 34, G: 139, B: 34, A: 255}
	}

	texteScore := canvas.NewText(fmt.Sprintf("%d / %d", score, total), couleurScore)
	texteScore.TextSize = 40
	texteScore.TextStyle = fyne.TextStyle{Bold: true}
	texteScore.Alignment = fyne.TextAlignCenter

	var texteMessage string
	if valide {
		texteMessage = fmt.Sprintf("Bravo %s ! Vous avez valide le theme \"%s\".", aq.nomJoueur, t.Nom)
	} else {
		texteMessage = fmt.Sprintf("Dommage %s... il faut au moins 3/5. Vous pouvez recommencer !", aq.nomJoueur)
	}

	message := widget.NewLabel(texteMessage)
	message.Alignment = fyne.TextAlignCenter
	message.Wrapping = fyne.TextWrapWord

	boutonRecommencer := widget.NewButton("Recommencer ce theme", func() {
		aq.afficherEcranQuiz(t)
	})
	if valide {
		boutonRecommencer.Disable()
	}

	boutonMenu := widget.NewButton("Retour au menu", func() {
		if len(aq.themesValides) == len(aq.themes) && len(aq.themes) > 0 {
			dialog.ShowInformation(
				"Felicitations !",
				fmt.Sprintf("Bravo %s ! Vous avez valide TOUS les themes !", aq.nomJoueur),
				aq.fenetre,
			)
		}
		aq.afficherEcranThemes()
	})
	boutonMenu.Importance = widget.HighImportance

	boutons := container.NewVBox(
		widget.NewSeparator(),
		boutonRecommencer,
		boutonMenu,
	)

	corps := container.NewVBox(
		titre,
		container.NewCenter(texteScore),
		message,
	)

	defilement := container.NewVScroll(corps)

	contenu := container.NewBorder(
		nil,
		boutons,
		nil,
		nil,
		defilement,
	)

	aq.fenetre.SetContent(container.NewPadded(contenu))
}