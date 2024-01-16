package src

import (
	"Goh-Lanta/temps"
	"net/http"
)

func Display(w http.ResponseWriter, r *http.Request) {
	data := GetAllAdv()
	temps.Temp.ExecuteTemplate(w, "display", data)
}
