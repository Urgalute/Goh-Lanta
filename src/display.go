package src

import (
	"Goh-Lanta/temps"
	"net/http"
)
//Affichage de tout les aventurier
func Display(w http.ResponseWriter, r *http.Request) {
	data := GetAllAdv()
	temps.Temp.ExecuteTemplate(w, "display", data)
}
