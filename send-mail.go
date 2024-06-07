package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func envoyerMailCommande(cmd Commande) {
	client := getClientByID(cmd.ClientID)
	produit := getProduitByID(cmd.ProduitID)

	from := "houenoujohannes60@gmail.com"
	to := client.Email
	body := fmt.Sprintf("Bonjour %s, votre commande de %d %s est confirmée. Prix: %.2f",
		client.Prenom, cmd.Quantite, produit.Titre, cmd.Prix)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Confirmation de commande\n\n" + body
		

	err := smtp.SendMail("localhost:1025",
		smtp.PlainAuth("", "", "", "localhost"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("Erreur d'envoi de l'email: %v\n", err)
	} else {
		fmt.Println("Email envoyé avec succès")
	}
}

func getClientByID(id int) Client {
	for _, c := range clients {
		if c.ID == id {
			return c
		}
	}
	return Client{}
}

func getProduitByID(id int) Produit {
	for _, p := range produits {
		if p.ID == id {
			return p
		}
	}
	return Produit{}
}
