package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
)

type Client struct {
	ID        int
	Nom       string
	Prenom    string
	Telephone string
	Adresse   string
	Email     string
}

var clients []Client

func addClient() {

	reader := bufio.NewReader(os.Stdin)

    var c Client
    c.ID = len(clients) + 1

    fmt.Print("Nom: ")
    c.Nom, _ = reader.ReadString('\n')

    fmt.Print("Prénom: ")
    c.Prenom, _ = reader.ReadString('\n')

    fmt.Print("Téléphone: ")
    c.Telephone, _ = reader.ReadString('\n')

    fmt.Print("Adresse: ")
    adresse, _ := reader.ReadString('\n')
    c.Adresse = strings.TrimSpace(adresse)

    fmt.Print("Email: ")
    email, _ := reader.ReadString('\n')
    c.Email = strings.TrimSpace(email)

    clients = append(clients, c)

	query := "INSERT INTO clients (nom, prenom, telephone, adresse, email) VALUES (?, ?, ?, ?, ?)"
    _, err := db.Exec(query, c.Nom, c.Prenom, c.Telephone, c.Adresse, c.Email)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Client ajouté avec succès")
}

func showClients() {
	for _, c := range clients {
		fmt.Printf("ID: %d, Nom: %s, Prénom: %s, Téléphone: %s, Adresse: %s, Email: %s\n", c.ID, c.Nom, c.Prenom, c.Telephone, c.Adresse, c.Email)
	}
}

func editClient() {
	var id int
	fmt.Print("ID du client à modifier: ")
	fmt.Scan(&id)

	for i, c := range clients {
		if c.ID == id {
			fmt.Print("Nom: ")
			fmt.Scan(&clients[i].Nom)
			fmt.Print("Prénom: ")
			fmt.Scan(&clients[i].Prenom)
			fmt.Print("Téléphone: ")
			fmt.Scan(&clients[i].Telephone)
			fmt.Print("Adresse: ")
			fmt.Scan(&clients[i].Adresse)
			fmt.Print("Email: ")
			fmt.Scan(&clients[i].Email)
			return
		}
	}
	fmt.Println("Client non trouvé")
}

func exportClientsCSV() {
	file, err := os.Create("clients.csv")
	if err != nil {
		log.Fatalln("Erreur de création de fichier:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Nom", "Prénom", "Téléphone", "Adresse", "Email"})
	for _, c := range clients {
		writer.Write([]string{
			fmt.Sprintf("%d", c.ID),
			c.Nom,
			c.Prenom,
			c.Telephone,
			c.Adresse,
			c.Email,
		})
	}
}
