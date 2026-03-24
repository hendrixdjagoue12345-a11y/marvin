package main

 
import "database/sql"
 
// on insere le joueur et on recupere son id
func insertJoueur(db *sql.DB, nom string) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO joueurs (nom, date_connexion) VALUES (?, NOW())",nom,)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
 
// sauvegarde le score a la fin d'une partie
func insertPartie(db *sql.DB, joueurID int64, themeID, score int) error {
	_, err := db.Exec(
		"INSERT INTO parties (joueur_id, theme_id, score, date_partie) VALUES (?, ?, ?, NOW())",joueurID, themeID, score,)
	return err
}
 
