package routes

import (
	"forum/controllers"
	"net/http"
)

func Jwt() {
	http.HandleFunc("/jwt", controllers.GetJwt)
}
