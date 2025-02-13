package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"forum/config"
	"forum/utils"
)

// func Usersapi{
// 	http.HandlerFunc("/users",con)
// }

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Nickname   string    `json:"nickname"`
	Email      string    `json:"email"`
	Age        string    `json:"age"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"createdat"`
}

func CreatApi(table string, w http.ResponseWriter, r *http.Request) {
	// Validate the table name to prevent SQL injection
	validTables := map[string]bool{"users": true, "products": true} // Add valid table names here
	if !validTables[table] {
		utils.CreateResponseAndLogger(w, http.StatusBadRequest, nil, "Invalid table name")
		return
	}

	// Construct the SQL query
	var query string
	if table == "users" {
		query = "SELECT username, nickname, email, age FROM " + table + ";"
	} else {
		query = "SELECT * FROM " + table + ";"
	}

	// Execute the query
	rows, err := config.DB.Query(query)
	if err != nil {
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Failed to execute query")
		return
	}
	defer rows.Close()
	// Process the query result
	var results []map[string]interface{}
	for rows.Next() {
		var username, nickname, email string
		var age int
		if table == "users" {
			err = rows.Scan(&username, &nickname, &email, &age)
		} else {
			// For other tables, dynamically handle columns (example for demonstration)
			columns, _ := rows.Columns()
			values := make([]interface{}, len(columns))
			valuePtrs := make([]interface{}, len(columns))
			for i := range columns {
				valuePtrs[i] = &values[i]
			}
			err = rows.Scan(valuePtrs...)
			if err == nil {
				result := make(map[string]interface{})
				for i, col := range columns {
					result[col] = values[i]
				}
				results = append(results, result)
			}
		}
		if err != nil {
			utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Failed to scan row")
			return
		}
		if table == "users" {
			results = append(results, map[string]interface{}{
				"username": username,
				"nickname": nickname,
				"email":    email,
				"age":      age,
			})
		}
	}

	// Check for errors after iteration
	if err = rows.Err(); err != nil {
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Error during rows iteration")
		return
	}
	// Encode the results into JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Failed to encode JSON")
		return
	}
	fmt.Println(w, ": w")
}

func Usersapi(w http.ResponseWriter, r *http.Request) {
	CreatApi("users", w, r)
}
