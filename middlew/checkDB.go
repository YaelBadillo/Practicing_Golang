package middlew

import (
	"net/http"

	"github.com/YaelJBS/Practicing_Golang/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == false {
			http.Error(w, "Lost connection to MongoDB", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
