package main

import (
	"fmt"
	"strings"
)

// Fonction principale d'un thème
// Elle affiche toutes les questions et retourne felicitation vous validez ce thème si le score est suffisant
func jouerTheme(nomJoueur string, questions []question) bool {
	score := 0
	
	// parcours des de toutes les questions du thème
	for _, q := range questions {

		// Affichage de la question et des choix
		fmt.Printf("\n%s\n", q.Question)
		fmt.Printf("A. %s\n", q.ChoixA)
		fmt.Printf("B. %s\n", q.ChoixB)
		fmt.Printf("C. %s\n", q.ChoixC)
		fmt.Printf("D. %s\n", q.ChoixD)

		//saisie de la réponse utilisateur
		var reponse string
		fmt.Print("Votre réponse : ")
		fmt.Scanln(&reponse)

		// Vérification de la réponse (insensible à la casse)
		if strings.EqualFold(reponse, q.Reponse) {
			fmt.Println("Bonne réponse !")
			score++
		} else {
			fmt.Printf("Mauvaise réponse.")
		}
	}

	// Affichage du score final
	fmt.Printf(
		"\n%s, votre score est %d/%d\n",
		nomJoueur,
		score,
		len(questions),
	)	

	// Si le score est suffisant, on valide le thème
	if score >= 3 {
		fmt.Println("Félicitations ! Vous avez validé ce thème.")
		return true
	} else {
		fmt.Println("Score insuffisant (minimum 3/5). Vous devez recommencer ce thème.")
		return false
	}
}
