package main

import "database/sql"

// connecterDB ouvre et verifie la connexion a la base de donnees MySQL
func connecterDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:cesi@tcp(127.0.0.1:3306)/datagouvschema")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// chargerThemes recupere tous les themes depuis la base de donnees
func chargerThemes(db *sql.DB) ([]theme, error) {
	lignes, err := db.Query("SELECT id, nom FROM themes")
	if err != nil {
		return nil, err
	}
	defer lignes.Close()

	var themes []theme
	for lignes.Next() {
		var t theme
		if err := lignes.Scan(&t.ID, &t.Nom); err != nil {
			return nil, err
		}
		themes = append(themes, t)
	}
	return themes, lignes.Err()
}

// chargerQuestions recupere les questions d'un theme donne
func chargerQuestions(db *sql.DB, themeID int) ([]question, error) {
	lignes, err := db.Query(`
		SELECT id, theme_id, question, choixA, choixB, choixC, choixD, bonne
		FROM questions
		WHERE theme_id = ?`, themeID)
	if err != nil {
		return nil, err
	}
	defer lignes.Close()

	var questions []question
	for lignes.Next() {
		var q question
		if err := lignes.Scan(&q.ID, &q.ThemeID, &q.Question, &q.ChoixA, &q.ChoixB, &q.ChoixC, &q.ChoixD, &q.Reponse); err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	return questions, lignes.Err()
}

// chargerHistorique recupere les parties jouees par un joueur
func chargerHistorique(db *sql.DB, joueurID int64) ([]partieHistorique, error) {
	lignes, err := db.Query(`
		SELECT t.nom, p.score, p.date_partie
		FROM parties p
		JOIN themes t ON p.theme_id = t.id
		WHERE p.joueur_id = ?
		ORDER BY p.date_partie DESC`, joueurID)
	if err != nil {
		return nil, err
	}
	defer lignes.Close()

	var historique []partieHistorique
	for lignes.Next() {
		var h partieHistorique
		if err := lignes.Scan(&h.Theme, &h.Score, &h.Date); err != nil {
			return nil, err
		}
		historique = append(historique, h)
	}
	return historique, lignes.Err()
}

// rechercherJoueurParNom cherche un joueur existant par son pseudo et renvoie son id
// renvoie 0, nil si le joueur n'existe pas
func rechercherJoueurParNom(db *sql.DB, nom string) (int64, error) {
	var id int64
	err := db.QueryRow(
		"SELECT id FROM joueurs WHERE nom = ? ORDER BY date_connexion DESC LIMIT 1", nom,
	).Scan(&id)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return id, nil
}

// chargerThemesValides recupere les themes valides (score >= 3) pour un joueur
func chargerThemesValides(db *sql.DB, joueurID int64) (map[int]bool, error) {
	lignes, err := db.Query(`
		SELECT DISTINCT theme_id
		FROM parties
		WHERE joueur_id = ? AND score >= 3`, joueurID)
	if err != nil {
		return nil, err
	}
	defer lignes.Close()

	valides := make(map[int]bool)
	for lignes.Next() {
		var themeID int
		if err := lignes.Scan(&themeID); err != nil {
			return nil, err
		}
		valides[themeID] = true
	}
	return valides, lignes.Err()
}