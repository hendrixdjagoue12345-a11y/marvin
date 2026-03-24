package main

import (
	"database/sql" 
)

func connectDB() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:cesi@tcp(127.0.0.1:3306)/datagouvschema")
	if err != nil {
		
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func loadThemes(db *sql.DB) ([]theme, error) {

	rows, err := db.Query("SELECT id, nom FROM themes")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var themes []theme

	for rows.Next() {
		var t theme

		if err := rows.Scan(&t.ID, &t.Nom); err != nil {
			return nil, err
		}

		themes = append(themes, t)
	}
	     return themes, rows.Err()
}

func loadQuestions(db *sql.DB, themeID int) ([]question, error) {

	rows, err := db.Query(`
		SELECT id, theme_id, question, choixA, choixB, choixC, choixD, bonne
		FROM questions
		WHERE theme_id = ?`, themeID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var questions []question

	for rows.Next() {
		var q question

		if err := rows.Scan(&q.ID, &q.ThemeID, &q.Question, &q.ChoixA, &q.ChoixB, &q.ChoixC, &q.ChoixD, &q.Reponse); err != nil {
            return nil, err
          }

	     questions = append(questions, q)
	    }
		
	 return questions, rows.Err()
}
