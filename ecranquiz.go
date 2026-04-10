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

// afficherEcranQuiz lance le quiz pour un theme donne
func (aq *appQuiz) afficherEcranQuiz(t theme) {
	questions, err := chargerQuestions(aq.db, t.ID)
	if err != nil {
		dialog.ShowError(err, aq.fenetre)
		return
	}
	if len(questions) == 0 {
		dialog.ShowInformation("Theme vide", "Aucune question disponible pour ce theme.", aq.fenetre)
		return
	}

	var score int
	var indexCourant int

	// attendReponse = true : le joueur doit choisir et valider
	// attendReponse = false : le joueur a valide et doit passer a la suite
	attendReponse := true

	etiquetteTitre := creerTitre(fmt.Sprintf("Theme : %s", t.Nom))

	etiquetteCompteur := widget.NewLabel("")
	etiquetteCompteur.Alignment = fyne.TextAlignCenter

	etiquetteQuestion := widget.NewLabel("")
	etiquetteQuestion.Wrapping = fyne.TextWrapWord
	etiquetteQuestion.TextStyle = fyne.TextStyle{Bold: true}

	etiquetteRetour := widget.NewLabel("")
	etiquetteRetour.Alignment = fyne.TextAlignCenter

	groupeChoix := widget.NewRadioGroup([]string{}, nil)

	var boutonValider *widget.Button

	// afficherQuestion affiche la question courante ou passe aux resultats
	var afficherQuestion func()
	afficherQuestion = func() {
		if indexCourant >= len(questions) {
			aq.afficherEcranResultat(t, score, len(questions))
			return
		}

		attendReponse = true
		q := questions[indexCourant]
		etiquetteCompteur.SetText(fmt.Sprintf("Question %d / %d", indexCourant+1, len(questions)))
		etiquetteQuestion.SetText(q.Question)
		etiquetteRetour.SetText("")

		groupeChoix.Options = []string{
			fmt.Sprintf("A.  %s", q.ChoixA),
			fmt.Sprintf("B.  %s", q.ChoixB),
			fmt.Sprintf("C.  %s", q.ChoixC),
			fmt.Sprintf("D.  %s", q.ChoixD),
		}
		groupeChoix.SetSelected("")
		groupeChoix.Refresh()
		boutonValider.SetText("Valider")
	}

	boutonValider = widget.NewButton("Valider", nil)
	boutonValider.Importance = widget.HighImportance

	boutonValider.OnTapped = func() {
		if attendReponse {
			choix := groupeChoix.Selected
			if choix == "" {
				etiquetteRetour.SetText("Veuillez choisir une reponse !")
				return
			}

			lettre := string([]rune(choix)[0])
			q := questions[indexCourant]

			if strings.EqualFold(lettre, q.Reponse) {
				etiquetteRetour.SetText("Bonne reponse !")
				score++
			} else {
				etiquetteRetour.SetText(fmt.Sprintf(
					"Mauvaise reponse. La bonne reponse etait : %s",
					strings.ToUpper(q.Reponse),
				))
			}

			indexCourant++
			attendReponse = false

			if indexCourant >= len(questions) {
				boutonValider.SetText("Voir les resultats")
			} else {
				boutonValider.SetText("Question suivante")
			}
		} else {
			afficherQuestion()
		}
	}

	boutonRetour := widget.NewButton("Retour aux themes", func() {
		dialog.ShowConfirm("Abandonner", "Voulez-vous abandonner ce theme ?", func(ok bool) {
			if ok {
				aq.afficherEcranThemes()
			}
		}, aq.fenetre)
	})

	contenu := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(etiquetteTitre),
		container.NewCenter(etiquetteCompteur),
		widget.NewSeparator(),
		etiquetteQuestion,
		groupeChoix,
		etiquetteRetour,
		boutonValider,
		widget.NewSeparator(),
		boutonRetour,
		layout.NewSpacer(),
	)

	aq.fenetre.SetContent(container.NewPadded(contenu))
	afficherQuestion()
}