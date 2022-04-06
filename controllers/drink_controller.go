package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetDrinks(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	idDrink := r.URL.Query()["id_drink"]
	drinkName := r.URL.Query()["name"]

	if idDrink == nil {
		query := "SELECT * FROM drinks WHERE name LIKE '" + drinkName[0] + "%'"
		rows, err := db.Query(query)
		CheckError(err)

		var drink Drink
		var drinks []Drink
		for rows.Next() {
			if err := rows.Scan(&drink.Id_Drink, &drink.Name, &drink.Price, &drink.Description); err != nil {
				PrintError(400, "No User Data Inserted To []User", w)
				log.Fatal(err)
				return
			} else {
				drinks = append(drinks, drink)
			}
		}

		var response DrinkResponse
		if err == nil {
			response.Status = 200
			response.Message = "Get Drinks Success"
			response.Data = drinks
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			PrintError(400, "No Data Found", w)
			return
		}
	} else {
		query := "SELECT * FROM drinks WHERE id_drink = " + idDrink[0]
		rows, err := db.Query(query)
		CheckError(err)

		var drink Drink
		var drinks []Drink
		for rows.Next() {
			if err := rows.Scan(&drink.Id_Drink, &drink.Name, &drink.Price, &drink.Description); err != nil {
				PrintError(400, "No User Data Inserted To []User", w)
				log.Fatal(err)
				return
			} else {
				drinks = append(drinks, drink)
			}
		}

		var response DrinkResponse
		if err == nil {
			response.Status = 200
			response.Message = "Get Drinks Success"
			response.Data = drinks
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			PrintError(400, "No Data Found", w)
			return
		}
	}
}
func AddDrinks(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		return
	}
	name := r.Form.Get("name")
	price := r.Form.Get("price")
	description := r.Form.Get("description")

	_, errQuery := db.Exec("INSERT INTO drinks (name, price, description) values (?,?,?)", name, price, description)

	if errQuery == nil {
		PrintSuccess(200, "Drink Inserted", w)
	} else {
		PrintError(400, "insert Drinks Failed", w)
	}
}
func DeleteDrink(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()
	idDrink := r.URL.Query()["id_drink"]

	_, errQuery := db.Exec("DELETE FROM drinks WHERE id_drink = ?", idDrink[0])

	if errQuery != nil {
		PrintSuccess(200, "Drink Deleted", w)
	} else {
		PrintError(400, "Delete Drink Failed", w)
		return
	}

}
func UpdateDrink(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()
	err := r.ParseForm()
	CheckError(err)
	idDrink := r.Form.Get("id_drink")
	name := r.Form.Get("name")
	price := r.Form.Get("price")
	description := r.Form.Get("description")

	_, errQuery := db.Exec("UPDATE drinks SET name=?, price=?, description=? WHERE id_drink=?", name, price, description, idDrink)

	if errQuery == nil {
		PrintSuccess(200, "Drink Updated", w)
	} else {
		PrintError(400, "Update Drinks Failed", w)
	}

}
