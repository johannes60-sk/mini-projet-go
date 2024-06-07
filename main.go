package main

import (
	"fmt"
)

func main() {

	initDB()

	for {
		fmt.Println("1. Ajouter un produit")
		fmt.Println("2. Afficher les produits")
		fmt.Println("3. Modifier un produit")
		fmt.Println("4. Supprimer un produit")
		fmt.Println("5. Exporter les produits en CSV")
		fmt.Println("6. Ajouter un client")
		fmt.Println("7. Afficher les clients")
		fmt.Println("8. Modifier un client")
		fmt.Println("9. Exporter les clients en CSV")
		fmt.Println("10. Effectuer une commande")
		fmt.Println("11. Exporter les commandes")
		fmt.Println("12. Quitter")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			addProduct()
		case 2:
			showProduct()
		case 3:
			editProduct()
		case 4:
			deleteProduct()
		case 5:
			exportProductCSV()
		case 6:
			addClient()
		case 7:
			showClients()
		case 8:
			editClient()
		case 9:
			exportClientsCSV()
		case 10:
			passCommande()
		case 11:
			exportCommandes()
		case 12:
			return
		default:
			fmt.Println("Choix invalide")
		}

	}
}
