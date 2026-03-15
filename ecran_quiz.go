package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (qa *quizApp) showQuizScreen(t theme) {
	questions, err := loadQuestions(qa.db, t.ID)
	if err != nil {
		dialog.ShowError(err, qa.window)
		return
	}
	if len(questions) == 0 {
		dialog.ShowInformation("Thème vide", "Aucune question disponible pour ce thème.", qa.window)
		return
	}

	var score int
	var currentIdx int
	// attendReponse = true quand on attend que le joueur valide
	// attendReponse = false quand le joueur a validé et doit passer à la suite
	attendReponse := true

	titleLabel := makeTitle(fmt.Sprintf("Thème : %s", t.Nom))

	counterLabel := widget.NewLabel("")
	counterLabel.Alignment = fyne.TextAlignCenter

	questionLabel := widget.NewLabel("")
	questionLabel.Wrapping = fyne.TextWrapWord
	questionLabel.TextStyle = fyne.TextStyle{Bold: true}

	feedbackLabel := widget.NewLabel("")
	feedbackLabel.Alignment = fyne.TextAlignCenter

	choixGroup := widget.NewRadioGroup([]string{}, nil)

	var validateBtn *widget.Button

	// Affiche la question courante
	var showQuestion func()
	showQuestion = func() {
		if currentIdx >= len(questions) {
			qa.showResultScreen(t, score, len(questions))
			return
		}

		attendReponse = true
		q := questions[currentIdx]
		counterLabel.SetText(fmt.Sprintf("Question %d / %d", currentIdx+1, len(questions)))
		questionLabel.SetText(q.Question)
		feedbackLabel.SetText("") // effacer le feedback

		choixGroup.Options = []string{
			fmt.Sprintf("A.  %s", q.ChoixA),
			fmt.Sprintf("B.  %s", q.ChoixB),
			fmt.Sprintf("C.  %s", q.ChoixC),
			fmt.Sprintf("D.  %s", q.ChoixD),
		}
		choixGroup.SetSelected("")
		choixGroup.Refresh()
		validateBtn.SetText("Valider →")
	}

	validateBtn = widget.NewButton("Valider →", nil)
	validateBtn.Importance = widget.HighImportance

	validateBtn.OnTapped = func() {
		if attendReponse {
			// — PHASE 1 : le joueur vient de choisir, on vérifie la réponse —
			selected := choixGroup.Selected
			if selected == "" {
				feedbackLabel.SetText("  Choisis une réponse !")
				return
			}

			lettre := string([]rune(selected)[0])
			q := questions[currentIdx]

			if strings.EqualFold(lettre, q.Reponse) {
				feedbackLabel.SetText("  Bonne réponse !")
				score++
			} else {
				feedbackLabel.SetText(fmt.Sprintf(
					"  Mauvaise réponse. La bonne réponse était : %s",
					strings.ToUpper(q.Reponse),
				))
			}

			currentIdx++
			attendReponse = false // on attend maintenant que le joueur passe à la suite

			if currentIdx >= len(questions) {
				validateBtn.SetText("Voir les résultats →")
			} else {
				validateBtn.SetText("Question suivante →")
			}

		} else {
			// — PHASE 2 : le joueur veut passer à la suite —
			showQuestion()
		}
	}

	backBtn := widget.NewButton("← Retour aux thèmes", func() {
		dialog.ShowConfirm("Abandonner", "Veux-tu abandonner ce thème ?", func(ok bool) {
			if ok {
				qa.showThemeScreen()
			}
		}, qa.window)
	})

	content := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(titleLabel),
		container.NewCenter(counterLabel),
		widget.NewSeparator(),
		questionLabel,
		choixGroup,
		feedbackLabel,
		validateBtn,
		widget.NewSeparator(),
		backBtn,
		layout.NewSpacer(),
	)

	qa.window.SetContent(container.NewPadded(content))
	showQuestion()
}