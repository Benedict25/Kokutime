package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func SeeCart(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	Id_Drink := r.Form.Get("Id_Drink")

	query := `SELECT drinks.name, drinks.price, detailed_carts.quantity 
	FROM detailed_carts JOIN drinks ON detailed_carts.id_drink=drinks.id_drink
	JOIN carts ON carts.id_cart=detailed_carts.id_cart
	JOIN users ON users.id_user=carts.id_user
	WHERE users.id_user = ?`

	if len(Id_Drink) > 0 {
		query += ` WHERE detailed_carts.Id_Drink = ` + Id_Drink
	}

	rows, err := db.Query(query, onlineId)

	if err != nil {
		log.Println(err)
		PrintError(400, "Rows Are Empty - Carts", w)
		return
	}

	var detailCart DetailCart
	var detailCarts []DetailCart
	for rows.Next() {
		if err := rows.Scan(&detailCart.DrinkData.Name,
			&detailCart.DrinkData.Price,
			&detailCart.Quantity); err != nil {
			log.Fatal(err.Error())
			PrintError(400, "No Product Data Inserted To []Product", w)
		} else {
			detailCarts = append(detailCarts, detailCart)
		}
	}

	var response DetailCartsResponse

	if len(detailCarts) > 0 {
		response.Data = detailCarts
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		PrintError(400, "No Detail Transaction In []detailTransactions", w)
		return
	}
}