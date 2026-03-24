package main
 
 
import (
	"database/sql"
 
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)
 
// quizApp c'est la structure principale de l'appli
type quizApp struct {
	fyneApp       fyne.App
	window        fyne.Window
	db            *sql.DB
	joueurID      int64
	nomJoueur     string
	themes        []theme      
	themesValides map[int]bool
}
 
// point d'entrée du programme
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
 