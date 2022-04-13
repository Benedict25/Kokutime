package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetPromoCode(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()
	idPromoCode := r.URL.Query()["id_promo"]

	query := "SELECT * FROM promos "

	if len(idPromoCode[0]) > 0 {
		query += "WHERE id_promo = " + idPromoCode[0]
	}
	row, err := db.Query(query)

	CheckError(err)

	var promo PromoCode
	var promos []PromoCode
	for row.Next() {
		row.Scan(
			&promo.Id_PromoCode,
			&promo.Minimal_Purchase,
			&promo.Promo_Amount,
			&promo.Quantity)
		if err != nil {
			log.Fatal(err.Error())
			PrintError(400, "Get PromoCode Failed", w)
			return
		} else {
			promos = append(promos, promo)
		}
	}

	var response PromoCodesResponse
	if err == nil {
		response.Status = 200
		response.Message = "Get PromoCode Success"
		response.Data = promos
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		PrintError(400, "No Data Found", w)
		return
	}

}
func AddPromoCode(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	CheckError(err)

	minimalPurchase := r.Form.Get("minimal_purchase")
	promoAmmount := r.Form.Get("promo_amount")
	quantity := r.Form.Get("quantity")

	_, errQuery := db.Exec("INSERT INTO promos (minimal_purchase, promo_amount, quantity) values (?,?,?)", minimalPurchase, promoAmmount, quantity)

	if errQuery == nil {
		PrintSuccess(200, "PromoCode Inserted", w)
	} else {
		PrintError(400, "insert PromoCode Failed", w)
	}
}
func DeletePromoCode(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()
	Id_PromoCode := r.URL.Query()["id_promo"]

	_, err := db.Exec("DELETE FROM promos WHERE id_promo = ?", Id_PromoCode[0])

	if err == nil {
		PrintSuccess(200, "PromoCode Deleted", w)
	} else {
		PrintError(400, "Delete PromoCode Failed", w)
		return
	}

}
func UpdatePromoCode(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()
	err := r.ParseForm()

	CheckError(err)

	idPromoCode := r.Form.Get("id_promo")
	minimalPurchase := r.Form.Get("minimal_purchase")
	promoAmmount := r.Form.Get("promo_amount")
	quantity := r.Form.Get("quantity")

	_, errQuery := db.Exec("UPDATE promos SET minimal_purchase=?, promo_amount=?, quantity=? WHERE id_promo=?", minimalPurchase, promoAmmount, quantity, idPromoCode)

	if errQuery == nil {
		PrintSuccess(200, "PromoCode Deleted", w)
	} else {
		PrintError(400, "Update PromoCode Failed", w)
	}

}
