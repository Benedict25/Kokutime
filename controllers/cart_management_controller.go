package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
func InsertCart(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	id_drink := r.Form.Get("id_drink")
	quantity := r.Form.Get("quantity")

	var userCart UserCart
	var userCarts []UserCart

	rows, err := db.Query(`SELECT detailed_carts.id_detailed_cart, detailed_carts.id_cart, detailed_carts.id_drink, detailed_carts.quantity 
	FROM detailed_carts 
	JOIN carts 
	ON detailed_carts.id_cart=carts.id_cart 
	WHERE carts.id_user =?`, onlineId)

	for rows.Next() {
		if err := rows.Scan(&userCart.idDetailedCart, &userCart.idCart, &userCart.idDrink, &userCart.quantity); err != nil {
			log.Fatal(err.Error())
			PrintError(400, "No Product Data Inserted To []Product", w)
		} else {
			userCarts = append(userCarts, userCart)
		}
	}

	id_drink_int, _ := strconv.Atoi(id_drink)
	isFound := false

	for i := 0; i < len(userCarts); i++ {
		if userCarts[i].idDrink == id_drink_int {
			_, errQuery := db.Exec("UPDATE detailed_carts SET quantity = quantity + "+quantity+" WHERE id_detailed_cart = ? ", userCarts[i].idDetailedCart)
			isFound = true
			if errQuery == nil {
				PrintSuccess(200, "Added To Cart", w)
			} else {
				PrintError(400, "Failed", w)
			}
			return
		}
	}
	if isFound != true {
		_, errQuery := db.Exec("INSERT INTO detailed_carts(id_cart, id_drink, quantity) VALUES (?,?,?)", userCarts[0].idCart, id_drink, quantity)

		if errQuery == nil {
			PrintSuccess(200, "Added To Cart", w)
		} else {
			PrintError(400, "Failed", w)
		}
	}

}
func UpdateQuantity(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	id_drink := r.Form.Get("id_drink")
	quantity := r.Form.Get("quantity")
	_, errQuery := db.Exec(`UPDATE detailed_carts 
	JOIN carts 
	ON detailed_carts.id_cart = carts.id_cart 
	SET detailed_carts.quantity = detailed_carts.quantity + `+quantity+` 
	WHERE detailed_carts.id_drink =? AND carts.id_user =?`, id_drink, onlineId)

	if errQuery == nil {
		PrintSuccess(200, "Updated Quantity", w)
	} else {
		PrintError(400, "Update Failed", w)
	}
}
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	id_drink := r.URL.Query()["id_drink"]

	_, errQuery := db.Exec(`DELETE detailed_carts
	FROM detailed_carts 
	JOIN carts
	ON detailed_carts.id_cart = carts.id_cart 
	WHERE detailed_carts.id_drink = ?
	AND carts.id_user =?`, id_drink[0], onlineId)

	if errQuery == nil {
		PrintSuccess(200, "Delete Success", w)
	} else {
		PrintError(400, "Delete Failed", w)
	}

}
