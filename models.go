package main

// structure du programme

// struture représentant une question du quiz
// elle correspond à la table 'questions' dans la base de données
type question struct {
	ID       int  // identifiant unique de la question
	ThemeID  int  // identifiant du thème associé à la question
	Question string  // texte de la question
	ChoixA   string  // choix de réponse A
	ChoixB   string  // choix de réponse B
	ChoixC   string  // choix de réponse C
	ChoixD   string  // choix de réponse D
	Reponse  string  // réponse correcte
}

// structure représentant un thème du quiz
// elle correspond à la table 'themes' dans la base de données
type theme struct {
	ID  int    // identifiant unique du thème
	Nom string // nom du thème
}


