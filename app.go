package main
 
 
import (
	"database/sql"
 
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)
 
// quizApp regroupe toutes les données partagées entre les écrans
type quizApp struct {
	fyneApp       fyne.App
	window        fyne.Window
	db            *sql.DB
	joueurID      int64
	nomJoueur     string
	themes        []theme      // type défini dans models.go
	themesValides map[int]bool
}
 
// main() — point d'entrée Fyne
func main() {
	a := app.New()
	w := a.NewWindow("Quiz Go · Fyne")
	w.Resize(fyne.NewSize(720, 520))
	w.SetFixedSize(true)
 
	qa := &quizApp{
		fyneApp:       a,
		window:        w,
		themesValides: make(map[int]bool),
	}
 
	qa.showLoginScreen()
	w.ShowAndRun()
}
 