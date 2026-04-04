package main
 
import (
	"database/sql"
	
)
 
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
 