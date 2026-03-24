package main
 
import (
	"fmt"
	"image/color"
	"strings"
 
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)
 // titre utilisé sur tous les ecrans
func makeTitle(text string) *canvas.Text {
	t := canvas.NewText(text, color.NRGBA{R: 0, G: 120, B: 215, A: 255})
	t.TextSize = 26
	t.TextStyle = fyne.TextStyle{Bold: true}
	t.Alignment = fyne.TextAlignCenter
	return t
}
 
func (qa *quizApp) showLoginScreen() {
	title := makeTitle(" Quiz Go")
 
	sub := canvas.NewText("Connecte-toi pour commencer", color.NRGBA{R: 80, G: 80, B: 80, A: 255})
	sub.Alignment = fyne.TextAlignCenter
 
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Ton prénom ou pseudo...")
 
	statusLabel := widget.NewLabel("")
	statusLabel.Alignment = fyne.TextAlignCenter
 
	connectBtn := widget.NewButton("Se connecter", func() {
		nom := strings.TrimSpace(nameEntry.Text)
		if nom == "" {
			statusLabel.SetText("  Entre ton nom pour continuer.")
			return
		}
 
		statusLabel.SetText("Connexion en cours...")
 
		// connexion a la bdd
		db, err := connectDB() 
		if err != nil {
			dialog.ShowError(fmt.Errorf("Impossible de joindre MySQL :\n%v", err), qa.window)
			statusLabel.SetText("")
			return
		}
 
		joueurID, err := insertJoueur(db, nom) 
		if err != nil {
			dialog.ShowError(fmt.Errorf("Erreur insertion joueur :\n%v", err), qa.window)
			return
		}
 
		themes, err := loadThemes(db) 
		if err != nil {
			dialog.ShowError(fmt.Errorf("Erreur chargement thèmes :\n%v", err), qa.window)
			return
		}
 
		// on stocke tout dans la structure
		qa.db = db
		qa.joueurID = joueurID
		qa.nomJoueur = nom
		qa.themes = themes
		qa.themesValides = make(map[int]bool)
 
		// ecran suivant
		qa.showThemeScreen()
	})
	connectBtn.Importance = widget.HighImportance
 
	content := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(title),
		container.NewCenter(sub),
		widget.NewSeparator(),
		widget.NewLabel(""),
		nameEntry,
		connectBtn,
		statusLabel,
		layout.NewSpacer(),
	)
 
	qa.window.SetContent(container.NewPadded(content))
}
 