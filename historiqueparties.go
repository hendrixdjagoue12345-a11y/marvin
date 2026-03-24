package main
 
import (
	"database/sql"
	"fmt"
)
 
// historique console
func showHistory(db *sql.DB, joueurID int64) error {
	rows, err := db.Query(
		"SELECT score, date_partie FROM parties WHERE joueur_id = ? ORDER BY date_partie DESC",
		joueurID,
	)
	if err != nil {
		return err
	}
	defer rows.Close()
 
	fmt.Println("\nHistorique des parties :")
	for rows.Next() {
		var score int
		var date string
		err := rows.Scan(&score, &date)
		if err != nil {
			return err
		}
		fmt.Printf("- %s | Score : %d\n", date, score)
	}
 
	return rows.Err()
}
 
// structure historique

type historiquePartie struct {
	Theme string 
	Score int    
	Date  string 
}
 
// historique fyne

func loadHistorique(db *sql.DB, joueurID int64) ([]historiquePartie, error) {
	rows, err := db.Query(`
		SELECT t.nom, p.score, p.date_partie
		FROM parties p
		JOIN themes t ON p.theme_id = t.id
		WHERE p.joueur_id = ?
		ORDER BY p.date_partie DESC`, joueurID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
 
	var hist []historiquePartie
	for rows.Next() {
		var h historiquePartie
		if err := rows.Scan(&h.Theme, &h.Score, &h.Date); err != nil {
			return nil, err
		}
		hist = append(hist, h)
	}
	return hist, rows.Err()
}
 