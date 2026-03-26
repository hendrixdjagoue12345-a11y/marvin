package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func generatePDF(history []historiquePartie, nomJoueur string) error {
	// ARIAL SIMPLE - PLUS DE POLICES EXTERNES
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(15, 15, 15)
	pdf.SetAutoPageBreak(true, 15)
	pdf.AddPage()

	// Titre
	pdf.SetFont("Arial", "B", 18)
	pdf.SetTextColor(30, 30, 30)
	pdf.CellFormat(0, 12, "Historique des parties", "", 1, "C", false, 0, "")
	
	// Infos (sans accents)
	pdf.SetFont("Arial", "", 11)
	pdf.SetTextColor(100, 100, 100)
	info := fmt.Sprintf("Joueur : %s | Genere le : %s", nomJoueur, time.Now().Format("02/01/2006 15:04"))
	pdf.CellFormat(0, 8, info, "", 1, "C", false, 0, "")

	// Entête tableau (sans accents)
	pdf.SetFont("Arial", "B", 12)
	pdf.SetFillColor(230, 230, 230)
	pdf.CellFormat(70, 10, "Theme", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 10, "Score", "1", 0, "C", true, 0, "")
	pdf.CellFormat(80, 10, "Date", "1", 1, "C", true, 0, "")

	// Lignes
	pdf.SetFont("Arial", "", 11)
	for i, h := range history {
		if i%2 == 0 {
			pdf.SetFillColor(250, 250, 250)
		} else {
			pdf.SetFillColor(240, 245, 255)
		}
		pdf.CellFormat(70, 10, h.Theme, "1", 0, "L", true, 0, "")
		
		scoreText := fmt.Sprintf("%d/5", h.Score)
		if h.Score >= 3 {
			pdf.SetTextColor(34, 139, 34)
		} else {
			pdf.SetTextColor(200, 50, 50)
		}
		pdf.CellFormat(30, 10, scoreText, "1", 0, "C", true, 0, "")
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(80, 10, h.Date, "1", 1, "L", true, 0, "")
	}

	pdf.Ln(5)
	pdf.SetFont("Arial", "I", 9)
	pdf.SetTextColor(120, 120, 120)
	pdf.CellFormat(0, 6, "Document genere automatiquement", "", 1, "R", false, 0, "")

	// Sauvegarde
	outputDir := "output"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}
	filename := fmt.Sprintf("historique_%s_%s.pdf", nomJoueur, time.Now().Format("20060102_150405"))
	fullPath := filepath.Join(outputDir, filename)
	return pdf.OutputFileAndClose(fullPath)
}
