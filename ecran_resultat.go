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

func (qa *quizApp) showResultScreen(t theme, score, total int) {
	_ = insertPartie(qa.db, qa.joueurID, t.ID, score)

	valide := score >= 3
	if valide {
		qa.themesValides[t.ID] = true
	}

	var titleTxt string
	if valide {
		titleTxt = "  Thème validé !"
	} else {
		titleTxt = "  Score insuffisant"
	}
	title := makeTitle(titleTxt)

	scoreColor := color.NRGBA{R: 200, G: 50, B: 50, A: 255}
	if valide {
		scoreColor = color.NRGBA{R: 34, G: 139, B: 34, A: 255}
	}
	scoreTxt := canvas.NewText(fmt.Sprintf("%d / %d", score, total), scoreColor)
	scoreTxt.TextSize = 40
	scoreTxt.TextStyle = fyne.TextStyle{Bold: true}
	scoreTxt.Alignment = fyne.TextAlignCenter

	var messageTxt string
	if valide {
		messageTxt = fmt.Sprintf("Bravo %s ! Tu as validé \"%s\".", qa.nomJoueur, t.Nom)
	} else {
		messageTxt = fmt.Sprintf("Dommage %s… il faut au moins 3/5. Tu peux recommencer !", qa.nomJoueur)
	}
	msg := widget.NewLabel(messageTxt)
	msg.Alignment = fyne.TextAlignCenter
	msg.Wrapping = fyne.TextWrapWord

	retryBtn := widget.NewButton("  Recommencer ce thème", func() {
		qa.showQuizScreen(t)
	})
	if valide {
		retryBtn.Disable()
	}

	menuBtn := widget.NewButton("  Retour au menu", func() {
		if len(qa.themesValides) == len(qa.themes) && len(qa.themes) > 0 {
			dialog.ShowInformation(
				" Félicitations !",
				fmt.Sprintf("Bravo %s ! Tu as validé TOUS les thèmes et remportes le Clavier d'Or !", qa.nomJoueur),
				qa.window,
			)
		}
		qa.showThemeScreen()
	})
	menuBtn.Importance = widget.HighImportance

	// boutons du bas
	boutons := container.NewVBox(
		widget.NewSeparator(),
		retryBtn,
		menuBtn,
	)

	body  := container.NewVBox(
		title,
		container.NewCenter(scoreTxt),
		msg,
	)
	scroll := container.NewVScroll(body)

	content := container.NewBorder(
		nil,
		boutons,
		nil,
		nil,
		scroll,
	)

	qa.window.SetContent(container.NewPadded(content))
}