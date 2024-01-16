package src

import (
	"Goh-Lanta/temps"
	"net/http"
	"strconv"
)
//Fonction qui affiche le template de la page pour mettre à jour les infos
func UpdatePage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("type"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	data := SearchId(id)
	temps.Temp.ExecuteTemplate(w, "updatepage", data)
}

//Fonction pour mettre à jour les infos
func Update(w http.ResponseWriter, r *http.Request) {











	http.Redirect(w, r, "/display", http.StatusSeeOther)
}