package main

import (
	"net/http"

	"github.com/Ricardolv/htmx-starter/handlers"
)

func main() {

	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("GET /", handlers.HTTPHandler(handlers.HomeHandler))
	http.Handle("GET /users", handlers.HTTPHandler(handlers.UsersHandler))
	http.Handle("GET /users/list", handlers.HTTPHandler(handlers.UsersListHandler))
	http.Handle("POST /stocks", handlers.HTTPHandler(handlers.StockHandler))

	http.ListenAndServe(":3000", nil)
}
