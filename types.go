package main

// question represente une question du quiz avec ses choix et la bonne reponse
type question struct {
	ID       int
	ThemeID  int
	Question string
	ChoixA   string
	ChoixB   string
	ChoixC   string
	ChoixD   string
	Reponse  string
}

// theme represente un theme disponible dans le quiz
type theme struct {
	ID  int
	Nom string
}

// partieHistorique contient les informations d'une partie jouee
type partieHistorique struct {
	Theme string
	Score int
	Date  string
}