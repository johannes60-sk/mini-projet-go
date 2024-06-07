package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
    "bufio"
    "strings"

)

type Produit struct {
	ID          int
	Titre       string
	Description string
	Prix        float64
	Quantite    int
	Actif       bool
}

var produits []Produit

func addProduct() {
	var p Produit
	p.ID = len(produits) + 1
	fmt.Print("Titre: ")
	fmt.Scan(&p.Titre)
	fmt.Print("Description: ")
	fmt.Scan(&p.Description)
	fmt.Print("Prix: ")
	fmt.Scan(&p.Prix)
	fmt.Print("Quantité: ")
	fmt.Scan(&p.Quantite)
	p.Actif = true
	produits = append(produits, p)

	query := "INSERT INTO produits (titre, description, prix, quantite, actif) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, p.Titre, p.Description, p.Prix, p.Quantite, p.Actif)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Produit ajouté avec succès")
}

func showProduct() {
	rows, err := db.Query("SELECT id, titre, description, prix, quantite, actif FROM produits")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var produits []Produit

	for rows.Next() {
		var p Produit
		err := rows.Scan(&p.ID, &p.Titre, &p.Description, &p.Prix, &p.Quantite, &p.Actif)
		if err != nil {
			log.Fatal(err)
		}
		produits = append(produits, p)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, p := range produits {
		fmt.Printf("ID: %d, Titre: %s, Description: %s, Prix: %.2f, Quantité: %d, Actif: %t\n", p.ID, p.Titre, p.Description, p.Prix, p.Quantite, p.Actif)
	}
}

func editProduct() {
	reader := bufio.NewReader(os.Stdin)

    var id int
    fmt.Print("Entrer l'id du produit à éditer: ")
    fmt.Scanf("%d", &id)

    reader.ReadString('\n')

    var p Produit
    fmt.Print("Nouveau Titre: ")
    p.Titre, _ = reader.ReadString('\n')
    p.Titre = strings.TrimSpace(p.Titre)

    fmt.Print("Nouvelle Description: ")
    p.Description, _ = reader.ReadString('\n')
    p.Description = strings.TrimSpace(p.Description)

    fmt.Print("Nouveau Prix: ")
    fmt.Scanf("%f", &p.Prix)

    fmt.Print("Nouvelle Quantité: ")
    fmt.Scanf("%d", &p.Quantite)

    query := "UPDATE produits SET titre = ?, description = ?, prix = ?, quantite = ? WHERE id = ?"
    _, err := db.Exec(query, p.Titre, p.Description, p.Prix, p.Quantite, id)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Produit mis à jour avec succès")

}

func deleteProduct() {
	var id int
	fmt.Print("Entrer l'id du produit a supprimer: ")
	fmt.Scanf("%d", &id)

	query := "DELETE FROM produits WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Produit supprimer avec succès")
}

func exportProductCSV() {
	file, err := os.Create("csv/produits.csv")
	if err != nil {
		log.Fatalln("Erreur de création de fichier:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Titre", "Description", "Prix", "Quantité", "Actif"})
	for _, p := range produits {
		writer.Write([]string{
			fmt.Sprintf("%d", p.ID),
			p.Titre,
			p.Description,
			fmt.Sprintf("%.2f", p.Prix),
			fmt.Sprintf("%d", p.Quantite),
			fmt.Sprintf("%t", p.Actif),
		})
	}
}
