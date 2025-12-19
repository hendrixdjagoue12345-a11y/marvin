// Programme principal du quiz
// Ce programmme permet à un utilisateur de jouer à un quiz en console
// Les questions et thèmes sont stockés dans une base de données MySQL


package main

// Imports des bibliothèques nécessaires
import (
	"database/sql" // permet d'interagir avec la base de données
	"fmt"          // permet l'affichage et la saisie en console
	"log"          // permet d'afficher les messages d'erreur
	"strings"      // permet de comparer les chaînes de caractères

	// Driver MySQL nécessaire pour se connecter à la base de données
	_ "github.com/go-sql-driver/mysql"

)




// Fonction principale du programme
func main() {
	// Connexion à la base de données
	db, err := sql.Open("mysql", "root:cesi@tcp(127.0.0.1:3306)/datagouvschema")
if err != nil {
	// Arrête le programme en cas d'erreur de connexion
	log.Fatal("Erreur connexion MySQL:", err)
}

// Ferme la connexion à la base de données à la fin du programme
defer db.Close()

// Vérifie que la connexion à la base de données est bien établie
if err := db.Ping(); err != nil {
 	log.Fatal("Impossible de ping MySQL:", err)	
}

fmt.Println("connexion MySQL réussie")

// demande du nom du joueur 
var nomJoueur string
fmt.Print("Entrez votre nom : ")
fmt.Scanln(&nomJoueur)

// chargement des thèmes depuis la base de données
themes, err := loadThemes(db)
if err != nil {
	log.Fatal(err)

}						

// Map permettant de stocker les thèmes déjà validés par le joueur
// La clé est l'ID du thème
themesValides := make(map[int]bool)

// Boucle principale du jeu 
for {

	// Affichage des thèmes disponibles
    fmt.Printf("\n%s, choisis un thème :\n", nomJoueur) 
	for _, t := range themes {
		// on affiche uniquement les thèmes non encore validés
		if !themesValides[t.ID] {
			fmt.Printf("%d - %s\n", t.ID, t.Nom)
		}
	}

	// Options pour quitter le jeu
	fmt.Println("0 - Quitter le jeu")

	var choixTheme int
	fmt.Print("quel thème veux-tu jouer ? ")
	fmt.Scanln(&choixTheme)

	// Si le joueur choisit de quitter le jeu
	if choixTheme == 0 {
		fmt.Println("Merci d'avoir joué ! Au revoir", nomJoueur)
		break
	}
	 
	// Vérifie que le thème choisi est valide
	if themesValides[choixTheme] {
		fmt.Println("Tu as déjà validé ce thème, choisis-en un autre.")
		continue
	}

	// Chargement des questions du thème choisi
	questions, err := loadQuestions(db, choixTheme)
	if err != nil {
		log.Fatal(err)
	}

	// verifie si il y a des questions pour le thème choisi
	if len(questions) == 0 {
		fmt.Println("Aucune question disponible pour ce thème.")
		continue
	}

	// tant que le score est insuffisant, on repose les questions
	for {
		score := 0
		
		// parcours des de toutes les questions du thème
		for _, q := range questions {

			// Affichage de la question et des choix
			fmt.Printf("\n%s\n", q.Question)
			fmt.Printf("A. %s\n", q.ChoixA)
			fmt.Printf("B. %s\n", q.ChoixB)
			fmt.Printf("C. %s\n", q.ChoixC)
			fmt.Printf("D. %s\n", q.ChoixD)

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
		fmt.Printf("\n%s, votre score pour le thème est %d/%d\n", nomJoueur, score, len(questions))

		// Si le score est suffisant, on valide le thème
		if score >= 3 {
			fmt.Println("Félicitations ! Vous avez validé ce thème.")
			themesValides[choixTheme] = true
			break
		} else {
			fmt.Println("Score insuffisant (minimum 3/5). tu dois recommencer ce thème.")
		}
	}		

	// Vérifie si tous les thèmes ont été validés
	if len(themesValides) == len(themes) {
		fmt.Printf("Félicitations %s ! Vous avez validé tous les thèmes du quiz et gagnez le clavier d'or.\n", nomJoueur)
		break
	}	
}
}	

// Fonction pour charger les thèmes depuis la base de données
func loadThemes(db *sql.DB) ([]theme, error) {

	// Requête SQL pour récupérer les thèmes
	rows, err := db.Query("SELECT id, nom FROM themes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var themes []theme

	// Parcours des résultats de la requête
	for rows.Next() {
		var t theme
		if err := rows.Scan(&t.ID, &t.Nom); err != nil {
			return nil, err
		}
		themes = append(themes, t)
	}

	return themes, rows.Err()
}

// Fonction pour charger les questions d'un thème depuis la base de données
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
        if err := rows.Scan(&q.ID, &q.ThemeID, &q.Question,
            &q.ChoixA, &q.ChoixB, &q.ChoixC, &q.ChoixD, &q.Reponse); err != nil {
            return nil, err
        }
        questions = append(questions, q)
    }
    return questions, rows.Err()
}
