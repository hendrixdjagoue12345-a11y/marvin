package main

 
import "database/sql"
 
// insertJoueur insère le joueur et retourne son ID
// (cette logique était directement dans mainConsole())
func insertJoueur(db *sql.DB, nom string) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO joueurs (nom, date_connexion) VALUES (?, NOW())",
		nom,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
 
// insertPartie enregistre le score d'une partie terminée
// (cette logique était dans la boucle de mainConsole())
func insertPartie(db *sql.DB, joueurID int64, themeID, score int) error {
	_, err := db.Exec(
		"INSERT INTO parties (joueur_id, theme_id, score, date_partie) VALUES (?, ?, ?, NOW())",
		joueurID, themeID, score,
	)
	return err
}
 
