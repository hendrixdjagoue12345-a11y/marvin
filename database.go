package main

import (
	"database/sql" // permet de gérer la connexion et les requêtes SQL
)


// Cette fonction établit la connexion à la base MySQL
// Elle retourne un pointeur vers sql.DB pour effectuer
// des requêtes sur la base et une erreur éventuelle
func connectDB() (*sql.DB, error) {

	// -------------------------------------------------
	// Ouverture de la connexion à la base
	// -------------------------------------------------
	// Format : utilisateur:motdepasse@tcp(ip:port)/nomBase
	// Ici on utilise 127.0.0.1 pour forcer le TCP/IP
	db, err := sql.Open("mysql", "root:cesi@tcp(127.0.0.1:3306)/datagouvschema")
	if err != nil {
		// Retourne nil et l'erreur si la connexion échoue
		return nil, err
	}

	// -------------------------------------------------
	// Vérification que la base est accessible
	// -------------------------------------------------
	// Ping permet de tester si la connexion est bien établie
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// -------------------------------------------------
	// Connexion réussie
	// -------------------------------------------------
	return db, nil
}

// -------------------------------------------------
// Fonction : loadThemes
// -------------------------------------------------
// Cette fonction récupère tous les thèmes disponibles
// depuis la table "themes" de la base de données
func loadThemes(db *sql.DB) ([]theme, error) {

	// -------------------------------------------------
	// Exécution de la requête SQL
	// -------------------------------------------------
	rows, err := db.Query("SELECT id, nom FROM themes")
	if err != nil {
		return nil, err
	}
	// Fermeture automatique des lignes à la fin de la fonction
	defer rows.Close()

	// -------------------------------------------------
	// Initialisation du slice pour stocker les thèmes
	// -------------------------------------------------
	var themes []theme

	// -------------------------------------------------
	// Parcours de chaque ligne retournée par la requête
	// -------------------------------------------------
	for rows.Next() {
		var t theme

		// Association des colonnes de la base aux champs de la structure
		if err := rows.Scan(&t.ID, &t.Nom); err != nil {
			return nil, err
		}

		// Ajout du thème au slice
		themes = append(themes, t)
	}

	// Retourne la liste des thèmes et une éventuelle erreur de parcours
	return themes, rows.Err()
}

// -------------------------------------------------
// Fonction : loadQuestions
// -------------------------------------------------
// Cette fonction récupère toutes les questions d'un thème donné
// depuis la table "questions" de la base de données
func loadQuestions(db *sql.DB, themeID int) ([]question, error) {

	// -------------------------------------------------
	// Exécution de la requête SQL avec paramètre
	// -------------------------------------------------
	// L'utilisation de "?" permet d'éviter les injections SQL
	rows, err := db.Query(`
		SELECT id, theme_id, question, choixA, choixB, choixC, choixD, bonne
		FROM questions
		WHERE theme_id = ?`, themeID)
	if err != nil {
		return nil, err
	}
	// Fermeture automatique des lignes à la fin de la fonction
	defer rows.Close()

	// -------------------------------------------------
	// Initialisation du slice pour stocker les questions
	// -------------------------------------------------
	var questions []question

	// -------------------------------------------------
	// Parcours de chaque ligne retournée par la requête
	// -------------------------------------------------
	for rows.Next() {
		var q question

		// Remplissage de la structure question avec les données de la ligne
		if err := rows.Scan(
			&q.ID,
			&q.ThemeID,
			&q.Question,
			&q.ChoixA,
			&q.ChoixB,
			&q.ChoixC,
			&q.ChoixD,
			&q.Reponse,
		); err != nil {
			return nil, err
		}

		// Ajout de la question au slice
		questions = append(questions, q)
	}

	// Retourne la liste des questions et une éventuelle erreur de parcours
	return questions, rows.Err()
}
