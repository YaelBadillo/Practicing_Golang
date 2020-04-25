package routers

import (
	"encoding/json"
	"net/http"

	"github.com/YaelJBS/Practicing_Golang/db"
	"github.com/YaelJBS/Practicing_Golang/models"
)

func ToPost(w http.ResponseWriter, r *http.Request) {
	var newPost models.Post

	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if len(newPost.Name) == 0 {
		http.Error(w, "Name required.", 400)
		return
	}

	if len(newPost.Text) == 0 {
		http.Error(w, "Text required.", 400)
		return
	}

	_, status, err := db.InsertPost(newPost)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Could not inserted Post.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
