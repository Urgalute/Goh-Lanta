package src

import (
	"Goh-Lanta/data"
	"Goh-Lanta/temps"
	"net/http"
	"strconv"
)

// Fonction qui affiche le template de l'aventurier pour mettre à jour les infos
func UpdatePage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("Update"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	data := SearchId(id)
	temps.Temp.ExecuteTemplate(w, "updatepage", data)
}

// Fonction qui supprime et rajoute un aventurier (update) en gardant le même ID
func Update(w http.ResponseWriter, r *http.Request) {
	AdvId, _ := strconv.Atoi(r.URL.Query().Get("type"))
	if AdvExist(int(AdvId)) {
		RemoveAdv(AdvId, true)
	}
	
	if r.Method == "POST" {
		var Adventurer data.AdvStruct
		Adventurer.Name = r.FormValue("name")
		Adventurer.Content = r.FormValue("content")
		Adventurer.Team = r.FormValue("team")
		Adventurer.Preview = r.FormValue("preview")
		//Adventurer.Img = handler.Filename
		Adventurer.Id = AdvId
		AddAdv(Adventurer, true)
	}
	http.Redirect(w, r, "/display", http.StatusSeeOther)
}
