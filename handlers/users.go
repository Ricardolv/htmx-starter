package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Ricardolv/htmx-starter/views"
)

func loadUsers() ([]string, error) {
	resp, err := http.Get("https://fakerapi.it/api/v1/users?_quantity=3")
	if err != nil {
		return nil, err
	}

	var user struct {
		Data []struct {
			FirstName string `json:"firstname"`
		} `json:"data"`
	}

	json.NewDecoder(resp.Body).Decode(&user)

	response := []string{}
	for _, u := range user.Data {
		response = append(response, u.FirstName)
	}

	return response, nil
}

func UsersHandler(w http.ResponseWriter, r *http.Request) error {
	users, err := loadUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	time.Sleep(1 * time.Second)
	return render(w, r, views.Users(users))
}

func UsersListHandler(w http.ResponseWriter, r *http.Request) error {
	users, err := loadUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return render(w, r, views.UsersList(users))
}
