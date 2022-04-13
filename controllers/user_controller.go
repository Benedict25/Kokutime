package controllers

import (
	"encoding/json"
	"net/http"
)

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM users WHERE id_user = ?", GetUserId(r))

	var user User
	err := row.Scan(&user.Id_User, &user.Name, &user.Address, &user.Email, &user.Password)

	var response UserResponse

	if err == nil {
		response.Status = 200
		response.Message = "Success"
		response.Data = user
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		PrintError(400, "User Not Found", w)
		return
	}
}

func EditProfile(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	CheckError(err)

	Name := r.Form.Get("name")
	Address := r.Form.Get("address")
	Email := r.Form.Get("email")
	Password := r.Form.Get("password")

	_, errQuery := db.Exec("UPDATE users SET name = ?, address = ?, email = ?, password = ? WHERE id_user = ?",
		Name,
		Address,
		Email,
		Password,
		GetUserId(r))

	if errQuery == nil {
		PrintSuccess(200, "Profile Updated", w)
	} else {
		PrintError(400, "Update Profile Failed", w)
		return
	}
}
