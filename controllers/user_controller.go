package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// DB + Query
	db := connect()
	defer db.Close()
	query := "SELECT * FROM users"

	// Get Data
	rows, err := db.Query(query)
	CheckError(err)

	// Insert Data To Array
	var user User
	var users []User
	for rows.Next() {
		if err := rows.Scan(&user.Id_User, &user.Name, &user.Address, &user.Email, &user.Password); err != nil {
			PrintError(400, "No User Data Inserted To []User", w)
			log.Fatal(err)
			return
		} else {
			users = append(users, user)
		}
	}

	// Show Result
	var response UsersResponse
	if len(users) > 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data = users
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		PrintError(400, "Users Not Found", w)
		return
	}
}
