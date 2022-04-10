package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetTotalPrice(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	row := db.QueryRow(`
	SELECT SUM(drinks.price * detailed_carts.quantity) 
	FROM drinks
	JOIN detailed_carts ON drinks.id_drink = detailed_carts.id_drink
	JOIN carts ON detailed_carts.id_cart = carts.id_cart
	WHERE carts.id_user = ?`,
		onlineId)

	var total int
	err := row.Scan(&total)

	if err == nil {
		PrintSuccess(200, "Total Price: "+strconv.Itoa(total), w)
	} else {
		PrintError(400, "User Not Found", w)
		return
	}
}

func CalculateTotalPrice(id_user int) int {
	db := connect()
	defer db.Close()

	row := db.QueryRow(`
	SELECT SUM(drinks.price * detailed_carts.quantity) 
	FROM drinks
	JOIN detailed_carts ON drinks.id_drink = detailed_carts.id_drink
	JOIN carts ON detailed_carts.id_cart = carts.id_cart
	WHERE carts.id_user = ?`, id_user)

	var total int
	err := row.Scan(&total)

	if err == nil {
		return total
	} else {
		return -1
	}
}

func GetMinimalPurchase(id_promo int) int {
	db := connect()
	defer db.Close()

	row := db.QueryRow(`SELECT minimal_purchase FROM promos WHERE id_promo = ?`, id_promo)

	var minimalPurchase int
	err := row.Scan(&minimalPurchase)

	row2 := db.QueryRow(`SELECT quantity FROM promos WHERE id_promo = ?`, id_promo) // Promo sold out

	var quantity int
	err2 := row2.Scan(&quantity)

	if quantity == 0 {
		return -1
	} else if err == nil && err2 == nil {
		return minimalPurchase
	} else {
		return -1
	}
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	currentTime := time.Now()

	err := r.ParseForm()
	CheckError(err)

	id_promo := r.Form.Get("id_promo")

	// Get data from detailed_carts
	rows, err := db.Query(`
	SELECT id_detailed_cart, detailed_carts.id_cart, id_drink, quantity 
	FROM detailed_carts
	JOIN carts ON detailed_carts.id_cart = carts.id_cart
	WHERE carts.id_user = ?`, onlineId)
	CheckError(err)

	var detailed_cart DetailedCart
	var detailed_carts []DetailedCart
	for rows.Next() {
		if err := rows.Scan(
			&detailed_cart.Id_Detailed_Cart,
			&detailed_cart.Id_Cart,
			&detailed_cart.Id_Drink,
			&detailed_cart.Quantity); err != nil {
			PrintError(400, "No Item In Cart", w)
			log.Fatal(err)
			return
		} else {
			detailed_carts = append(detailed_carts, detailed_cart)
		}
	}

	// Check if total price > minimal purchase
	totalPrice := CalculateTotalPrice(onlineId)

	var minimalPurchase int

	if len(id_promo) > 0 { // Inserted promo code in form
		id_promo_int, _ := strconv.Atoi(id_promo)
		minimalPurchase = GetMinimalPurchase(id_promo_int) // If promo not found / quantity = 0 -> return -1
	} else { // Doesn't use promo code
		minimalPurchase = 0 // Uses promo_code = 0
	}

	if totalPrice == -1 {
		PrintError(400, "No Item In Cart", w)
		return
	} else if minimalPurchase == -1 {
		PrintError(400, "Promo Code Not Found / Out Of Stock", w)
		return
	} else if totalPrice > minimalPurchase { // Checkout Success
		// Create new transaction
		if len(id_promo) > 0 { // Inserted promo code in form
			db.Exec(`
			INSERT INTO transactions(id_user, id_promo, status, date) 
			VALUES(?, ?, ?, ?)`, onlineId, id_promo, "Processing", currentTime.Format("2006-01-02"))

			db.Exec(`
			UPDATE promos 
			SET quantity = quantity - 1 
			WHERE id_promo = ?`, id_promo) // Reduce promo quantity
		} else {
			db.Exec(`
			INSERT INTO transactions(id_user, id_promo, status, date) 
			VALUES(?, ?, ?, ?)`, onlineId, 0, "Processing", "2002-04-01")
		}

		// Get newest id_transaction for the person
		rows, err := db.Query(`
		SELECT id_transaction 
		FROM transactions 
		WHERE id_user = ?`, onlineId)
		CheckError(err)

		var id_transaction int
		var id_transactions []int

		for rows.Next() {
			if err := rows.Scan(
				&id_transaction); err != nil {
				PrintError(400, "Error in Inserting Transaction", w)
				log.Fatal(err)
				return
			} else {
				id_transactions = append(id_transactions, id_transaction)
			}
		}

		// Insert into detailed_transactions
		for i := range detailed_carts {
			db.Exec(`INSERT INTO detailed_transactions(id_transaction, id_drink, quantity) 
			VALUES(?, ?, ?)`, id_transactions[len(id_transactions)-1], detailed_carts[i].Id_Drink, detailed_carts[i].Quantity)
		}

		// Delete from detailed_cart
		db.Exec(`DELETE detailed_carts
		FROM detailed_carts 
		JOIN carts ON detailed_carts.id_cart = carts.id_cart
		WHERE carts.id_user = ?`, onlineId)

		text := "Received payment: Rp." + strconv.Itoa(totalPrice)
		SendMail("cobapbp@gmail.com", "Thanks For Ordering", text)

		PrintSuccess(200, "Checked out", w)
	} else { // totalPrice < minimalPurchase
		PrintError(400, "Minimum Purchase Not Fulfilled", w)
		return
	}

}
