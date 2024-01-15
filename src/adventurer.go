package src

import (
	"Goh-Lanta/temps"
	"net/http"
)

func DisplayAdv(w http.ResponseWriter, r *http.Request) {
	temps.Temp.ExecuteTemplate(w, "adventurer", nil)
}
