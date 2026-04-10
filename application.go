package main

import (
	"database/sql"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	_ "github.com/go-sql-driver/mysql"
)

// appQuiz est la structure principale de l'application
type appQuiz struct {
	fyneApp       fyne.App
	fenetre       fyne.Window
	db            *sql.DB
	joueurID      int64
	nomJoueur     string
	themes        []theme
	themesValides map[int]bool
}

// point d'entree du programme
func main() {
	a := app.New()
	f := a.NewWindow("Quiz Go")
	f.Resize(fyne.NewSize(720, 520))
	f.SetFixedSize(true)

	aq := &appQuiz{
		fyneApp:       a,
		fenetre:       f,
		themesValides: make(map[int]bool),
	}

	aq.afficherEcranAccueil()
	f.ShowAndRun()
}