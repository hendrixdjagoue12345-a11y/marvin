package main

// structure du programme

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

// structure représentant un thème du quiz
// elle correspond à la table 'themes' dans la base de données
type theme struct {
	ID  int   
	Nom string 
}


