package main

import (
	"fmt"
	"time"
)

type Commande struct {
	ID        int
	ClientID  int
	ProduitID int
	Quantite  int
	Prix      float64
	DateAchat time.Time
}

var commandes []Commande

func passCommande() {
	var cmd Commande
	cmd.ID = len(commandes) + 1
	fmt.Print("ID du client: ")
	fmt.Scan(&cmd.ClientID)
	fmt.Print("ID du produit: ")
	fmt.Scan(&cmd.ProduitID)
	fmt.Print("Quantité: ")
	fmt.Scan(&cmd.Quantite)

	for _, p := range produits {
		if p.ID == cmd.ProduitID && p.Actif && p.Quantite >= cmd.Quantite {
			cmd.Prix = float64(cmd.Quantite) * p.Prix
			p.Quantite -= cmd.Quantite
			cmd.DateAchat = time.Now()
			commandes = append(commandes, cmd)
			envoyerMailCommande(cmd)
			genererPDFCommande(cmd)
			return
		}
	}
	fmt.Println("Produit non trouvé, inactif ou quantité insuffisante")
}
