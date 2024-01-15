package src

import (
	"Goh-Lanta/temps"
	"net/http"
)

func Display(w http.ResponseWriter, r *http.Request) {
	temps.Temp.ExecuteTemplate(w, "display", nil)
}
