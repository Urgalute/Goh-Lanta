package src

import (
	"Goh-Lanta/data"
	"Goh-Lanta/temps"
	"net/http"
	"strconv"
)

func DisplayAdv(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("type"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	data := SearchId(id)
	temps.Temp.ExecuteTemplate(w, "adventurer", data)
}

// recherche par ID
func SearchId(id int) []data.AdvStruct {
	var pertinentAdv []data.AdvStruct
	for _, adventurer := range Adventurer {
		if adventurer.Id == id {
			pertinentAdv = append(pertinentAdv, adventurer)
		}
	}
	return pertinentAdv
}
