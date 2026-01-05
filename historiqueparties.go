package main
import (
    "database/sql"
    "fmt"
)

func showHistory(db *sql.DB, joueurID int64) error {
    rows, err := db.Query(
        "SELECT score, date_partie FROM parties WHERE joueur_id = ? ORDER BY date_partie DESC",
        joueurID,
    )
    if err != nil {
        return err
    }
    defer rows.Close()

    fmt.Println("\nHistorique des parties :")
    for rows.Next() {
        var score int
        var date string
        err := rows.Scan(&score, &date)
        if err != nil {
            return err
        }
        fmt.Printf("- %s | Score : %d\n", date, score)
    }

    return rows.Err()
}
