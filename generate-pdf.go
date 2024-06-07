package main


import (
	"fmt"
	"log"
	"github.com/phpdave11/gofpdf"
	"os"
	"encoding/csv"
)

func genererPDFCommande(cmd Commande) {
    client := getClientByID(cmd.ClientID)
    produit := getProduitByID(cmd.ProduitID)

    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, "Facture")
    pdf.Ln(20)

    pdf.SetFont("Arial", "", 12)
    pdf.Cell(40, 10, fmt.Sprintf("Client: %s %s", client.Prenom, client.Nom))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Produit: %s", produit.Titre))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Quantité: %d", cmd.Quantite))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Prix: %.2f", cmd.Prix))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Date d'achat: %s", cmd.DateAchat.Format("02-01-2006 15:04:05")))

    err := pdf.OutputFileAndClose("commande.pdf")
    if err != nil {
        log.Printf("Erreur de génération de PDF: %v\n", err)
    } else {
        fmt.Println("PDF généré avec succès")
    }
}




func exportCommandes() {
    file, err := os.Create("commandes.csv")
    if err != nil {
        log.Fatalln("Erreur de création de fichier:", err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.Write([]string{"ID", "ClientID", "ProduitID", "Quantité", "Prix", "DateAchat"})
    for _, cmd := range commandes {
        writer.Write([]string{
            fmt.Sprintf("%d", cmd.ID),
            fmt.Sprintf("%d", cmd.ClientID),
            fmt.Sprintf("%d", cmd.ProduitID),
            fmt.Sprintf("%d", cmd.Quantite),
            fmt.Sprintf("%.2f", cmd.Prix),
            cmd.DateAchat.Format("02-01-2006 15:04:05"),
        })
    }
}
