package handlers

import (
	"math/rand"
	"net/http"

	"github.com/Ricardolv/htmx-starter/views"
)

func StockHandler(w http.ResponseWriter, r *http.Request) error {
	price := rand.Float64()
	return render(w, r, views.Stock(price))
}
