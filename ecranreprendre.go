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

// afficherEcranReprendre affiche le formulaire pour rechercher une partie existante
func (aq *appQuiz) afficherEcranReprendre() {
	titre := creerTitre("Reprendre une partie")

	sousTitre := canvas.NewText("Entrez votre pseudo pour retrouver votre progression", color.NRGBA{R: 80, G: 80, B: 80, A: 255})
	sousTitre.Alignment = 1 // centrer

	champNom := widget.NewEntry()
	champNom.SetPlaceHolder("Votre prenom ou pseudo...")

	etiquetteStatut := widget.NewLabel("")
	etiquetteStatut.Alignment = 1 // centrer

	boutonRechercher := widget.NewButton("Rechercher", func() {
		nom := strings.TrimSpace(champNom.Text)
		if nom == "" {
			etiquetteStatut.SetText("Veuillez entrer un nom pour rechercher.")
			return
		}

		etiquetteStatut.SetText("Recherche en cours...")

		db, err := connecterDB()
		if err != nil {
			dialog.ShowError(fmt.Errorf("Impossible de joindre la base de donnees :\n%v", err), aq.fenetre)
			etiquetteStatut.SetText("")
			return
		}

		// Recherche du joueur existant
		joueurID, err := rechercherJoueurParNom(db, nom)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Erreur lors de la recherche :\n%v", err), aq.fenetre)
			etiquetteStatut.SetText("")
			return
		}

		if joueurID == 0 {
			// Aucun joueur trouve avec ce pseudo
			etiquetteStatut.SetText(fmt.Sprintf("Aucune partie trouvee pour \"%s\".", nom))
			return
		}

		// Joueur trouve, on charge ses donnees
		themes, err := chargerThemes(db)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Erreur lors du chargement des themes :\n%v", err), aq.fenetre)
			return
		}

		themesValides, err := chargerThemesValides(db, joueurID)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Erreur lors du chargement de la progression :\n%v", err), aq.fenetre)
			return
		}

		aq.db = db
		aq.joueurID = joueurID
		aq.nomJoueur = nom
		aq.themes = themes
		aq.themesValides = themesValides

		// On cree une nouvelle entree de connexion pour tracer la reprise
		_, _ = insererJoueur(db, nom)

		dialog.ShowInformation(
			"Partie retrouvee",
			fmt.Sprintf("Bienvenue de retour, %s ! Votre progression a ete chargee.", nom),
			aq.fenetre,
		)

		aq.afficherEcranThemes()
	})
	boutonRechercher.Importance = widget.HighImportance

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
		boutonRechercher,
		etiquetteStatut,
		widget.NewSeparator(),
		boutonRetour,
		layout.NewSpacer(),
	)

	aq.fenetre.SetContent(container.NewPadded(contenu))
}