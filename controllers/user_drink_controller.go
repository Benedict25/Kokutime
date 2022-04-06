package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func SeeDrinks(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()
	query := "SELECT * FROM drinks"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		PrintError(400, "No Drink Data Inserted To []Drink", w)
	}

	var drink Drink
	var drinks []Drink
	for rows.Next() {
		if err := rows.Scan(&drink.Id_Drink, &drink.Name, &drink.Price, &drink.Description); err != nil {
			log.Fatal(err.Error())
			PrintError(400, "No Drink Data Inserted To []Drink", w)
		} else {
			drinks = append(drinks, drink)
		}
	}

	var response DrinksResponse

	if len(drinks) > 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data = drinks
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		PrintError(400, "null []Drink", w)
	}
}
func SeeDetailedDrinks(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()
	query := "SELECT * FROM drinks WHERE id_drink=?"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		PrintError(400, "No Detail Drink Data Inserted To []Detail Drink", w)
	}

	var drink Drink
	var drinks []Drink
	for rows.Next() {
		if err := rows.Scan(&drink.Id_Drink, &drink.Name, &drink.Price, &drink.Description); err != nil {
			log.Fatal(err.Error())
			PrintError(400, "No Detail Drink Data Inserted To []Detail Drink", w)
		} else {
			drinks = append(drinks, drink)
		}
	}

	var response DrinksResponse
	if len(drinks) > 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data = drinks
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		PrintError(400, "null []Detail Drink", w)
	}
}
