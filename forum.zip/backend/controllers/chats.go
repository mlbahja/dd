package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"forum/config"
	"forum/models"
	"forum/utils"
	"net/http"
	"strings"
)

/*
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := errors.New("method not allowed")
		utils.CreateResponseAndLogger(w, http.StatusMethodNotAllowed, err, "Method not allowed")
		return
	}
	var user models.User
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Internal server error")
		return
	}
	user.Username = html.EscapeString(user.Username)
	user.Email = html.EscapeString(user.Email)
	user.Password = html.EscapeString(user.Password)

	if err := utils.Validation(user, true); err != nil {
		utils.CreateResponseAndLogger(w, http.StatusBadRequest, err, err.Error())
		return
	}

	if err := utils.Hash(&user.Password); err != nil {
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Internal server error")
		return
	}

	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	_, err := config.DB.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.username" {
			utils.CreateResponseAndLogger(w, http.StatusBadRequest, err, "Username already exists")
			return
		} else if err.Error() == "UNIQUE constraint failed: users.email" {
			utils.CreateResponseAndLogger(w, http.StatusBadRequest, err,
				"Email already exists")
			return
		}
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err,
			"Internal server error")
		return
	}
	utils.CreateResponseAndLogger(w, http.StatusCreated, nil, "user created successfully")
}
*/
/*
	type Message struct {
		Message  string `json:"message"`
		Sender   string `json:"sender"`
		Receiver string `json:"receiver"`
		Created  string `json:"created"`
		Status   string `json:"status"`
	}

	type Chat struct {
		User1 string `json:"user1"`
		User2 string `json:"user2"`
	}
*/
func Chats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := errors.New("method not allowed")
		utils.CreateResponseAndLogger(w, http.StatusMethodNotAllowed, err, "Method not allowed")
		return
	}
	// var message models.Message
	var chat models.Chat
	var chatsBytes []byte

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&chat); err != nil {
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Internal server error")
		return
	}

	var readable []map[string]string

	query := "SELECT * FROM chats;"

	_, err := config.DB.Exec(query)
	if err != nil {
		utils.CreateResponseAndLogger(w, http.StatusBadRequest, err, "can't Exite query !!!")
		return
	}
	utils.CreateResponseAndLogger(w, http.StatusCreated, nil, "query are execute successfully ")

	chatsBytes = []byte(query)

	json.Unmarshal(chatsBytes, &readable)
	var listOFstrings []string
	for _, obj := range readable {
		listOFstrings = append(listOFstrings, obj["uer1"]+"~"+obj["user2"])
	}
	flag := 0
	for _, el := range listOFstrings {
		users := strings.Split(el, "~")
		if users[0] == chat.User1 && users[1] == chat.User2 || users[0] == chat.User2 && users[1] == chat.User1 {
			flag = 1
		}
	}
	fmt.Printf("\n")
	if flag == 0 {

		query := "INSERT INTO chats(user1, user2) values(?,?)"

		_, err := config.DB.Exec(query, chat.User1, chat.User2)
		if err != nil {
			utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err,
				"Internal server error")
			return
		}
		utils.CreateResponseAndLogger(w, http.StatusCreated, nil, "users inserted successfully")
	}
}
