package main

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// afficherEcranConnexion affiche le formulaire de saisie du pseudo pour une nouvelle partie
func (aq *appQuiz) afficherEcranConnexion() {
	titre := creerTitre("Nouvelle partie")

	sousTitre := canvas.NewText("Entrez votre pseudo pour commencer", color.NRGBA{R: 80, G: 80, B: 80, A: 255})
	sousTitre.Alignment = 1 // centrer

	champNom := widget.NewEntry()
	champNom.SetPlaceHolder("Votre prenom ou pseudo...")

	etiquetteStatut := widget.NewLabel("")
	etiquetteStatut.Alignment = 1 // centrer

	boutonConnexion := widget.NewButton("Commencer", func() {
		nom := strings.TrimSpace(champNom.Text)
		if nom == "" {
			etiquetteStatut.SetText("Veuillez entrer un nom pour continuer.")
			return
		}

		etiquetteStatut.SetText("Connexion en cours...")

		db, err := connecterDB()
		if err != nil {
			dialog.ShowError(fmt.Errorf("Impossible de joindre la base de donnees :\n%v", err), aq.fenetre)
			etiquetteStatut.SetText("")
			return
		}

		joueurID, err := insererJoueur(db, nom)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Erreur lors de la creation du joueur :\n%v", err), aq.fenetre)
			return
		}

		themes, err := chargerThemes(db)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Erreur lors du chargement des themes :\n%v", err), aq.fenetre)
			return
		}

		aq.db = db
		aq.joueurID = joueurID
		aq.nomJoueur = nom
		aq.themes = themes
		aq.themesValides = make(map[int]bool)

		aq.afficherEcranThemes()
	})
	boutonConnexion.Importance = widget.HighImportance

	boutonRetour := widget.NewButton("Retour", func() {
		aq.afficherEcranAccueil()
	})

	contenu := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(titre),
		container.NewCenter(sousTitre),
		widget.NewSeparator(),
		widget.NewLabel(""),
		champNom,
		boutonConnexion,
		etiquetteStatut,
		widget.NewSeparator(),
		boutonRetour,
		layout.NewSpacer(),
	)

	aq.fenetre.SetContent(container.NewPadded(contenu))
}