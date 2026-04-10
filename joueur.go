package main

import "database/sql"

// insererJoueur cree un nouveau joueur et renvoie son identifiant
func insererJoueur(db *sql.DB, nom string) (int64, error) {
	resultat, err := db.Exec(
		"INSERT INTO joueurs (nom, date_connexion) VALUES (?, NOW())", nom,
	)
	if err != nil {
		return 0, err
	}
	return resultat.LastInsertId()
}

// enregistrerPartie sauvegarde le score d'une partie terminee
func enregistrerPartie(db *sql.DB, joueurID int64, themeID, score int) error {
	_, err := db.Exec(
		"INSERT INTO parties (joueur_id, theme_id, score, date_partie) VALUES (?, ?, ?, NOW())",
		joueurID, themeID, score,
	)
	return err
}