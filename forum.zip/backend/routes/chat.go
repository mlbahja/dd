package routes

import (
	"forum/controllers"
	"net/http"
)

func WebSocket() {
	http.HandleFunc("/chat", controllers.Chats)
	// http.HandleFunc(w http.ResponseWriter, r *http.Request){
	// 	if r.URL.Path != "/chat"{
	// 		http.Error(w, "page not found",http.StatusFound)
	// 		return
	// 	}
	// 	var Message message
	// 	if err := json.NewDecoder(r.Body).Decode(&Message);err != nil {
	// 		http.Error(w, "Invalid Input", http.StatusBadRequest)
	// 		return
	// 	}
	// 	defer r.Body.Close()

	// }
	H
}
