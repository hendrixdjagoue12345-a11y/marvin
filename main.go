
package main
 
// Imports des bibliothèques nécessaires
import (
	"database/sql" 
	"fmt"          
	"log"        
	"strings"    
 
	
	_ "github.com/go-sql-driver/mysql"
)
 
func mainConsole() {
	
	db, err := sql.Open("mysql", "root:cesi@tcp(127.0.0.1:3306)/datagouvschema")
	if err != nil {
		log.Fatal("Erreur connexion MySQL:", err)
	}
	defer db.Close()
 
	if err := db.Ping(); err != nil {
		log.Fatal("Impossible de ping MySQL:", err)
	}
 
	fmt.Println("connexion MySQL réussie")
 
	var nomJoueur string
	fmt.Print("Entrez votre nom : ")
	fmt.Scanln(&nomJoueur)
 
	result, err := db.Exec(
		"INSERT INTO joueurs (nom,date_connexion) VALUES (?, NOW())",
		nomJoueur,
	)
	if err != nil {
		log.Fatal("Erreur insertion joueur:", err)
	}
 
	joueurID, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Erreur récupération ID joueur:", err)
	}
 
	themes, err := loadThemes(db)
	if err != nil {
		log.Fatal(err)
	}
 
	themesValides := make(map[int]bool)
 
	for {
		fmt.Printf("\n%s, choisis un thème :\n", nomJoueur)
		for _, t := range themes {
			if !themesValides[t.ID] {
				fmt.Printf("%d - %s\n", t.ID, t.Nom)
			}
		}
		fmt.Println("0 - Quitter le jeu")
 
		var choixTheme int
		fmt.Print("quel thème veux-tu jouer ? ")
		fmt.Scanln(&choixTheme)
 
		if choixTheme == 0 {
			fmt.Println("Merci d'avoir joué ! Au revoir", nomJoueur)
			break
		}
 
		if themesValides[choixTheme] {
			fmt.Println("Tu as déjà validé ce thème, choisis-en un autre.")
			continue
		}
 
		questions, err := loadQuestions(db, choixTheme)
		if err != nil {
			log.Fatal(err)
		}
 
		if len(questions) == 0 {
			fmt.Println("Aucune question disponible pour ce thème.")
			continue
		}
 
		for {
			score := 0
 
			for _, q := range questions {
				fmt.Printf("\n%s\n", q.Question)
				fmt.Printf("A. %s\n", q.ChoixA)
				fmt.Printf("B. %s\n", q.ChoixB)
				fmt.Printf("C. %s\n", q.ChoixC)
				fmt.Printf("D. %s\n", q.ChoixD)
 
				var reponse string
				fmt.Print("Votre réponse : ")
				fmt.Scanln(&reponse)
 
				if strings.EqualFold(reponse, q.Reponse) {
					fmt.Println("Bonne réponse !")
					score++
				} else {
					fmt.Printf("Mauvaise réponse.")
				}
			}
 
			fmt.Printf("\n%s, votre score pour le thème est %d/%d\n", nomJoueur, score, len(questions))
 
			_, err := db.Exec(
				"INSERT INTO parties (joueur_id, theme_id, score, date_partie) VALUES (?, ?, ?, NOW())",
				joueurID, choixTheme, score,
			)
			if err != nil {
				log.Fatal("Erreur insertion partie:", err)
			}
 
			if score >= 3 {
				fmt.Println("Félicitations ! Vous avez validé ce thème.")
				themesValides[choixTheme] = true
				break
			} else {
				fmt.Println("Score insuffisant (minimum 3/5). tu dois recommencer ce thème.")
			}
		}
 
		if len(themesValides) == len(themes) {
			fmt.Printf("Félicitations %s ! Vous avez validé tous les thèmes du quiz et gagnez le clavier d'or.\n", nomJoueur)
			break
		}
	}
}
 