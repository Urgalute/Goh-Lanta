package src

import (
	"Goh-Lanta/data"
	"net/http"
	"strconv"
)

// fonction pour supprimer un aventurier du tableau et du json
func Remove(w http.ResponseWriter, r *http.Request) {
	AdvId, _ := strconv.Atoi(r.FormValue("Suppr"))
	if AdvExist(int(AdvId)) {
		RemoveAdv(AdvId, true)
		http.Redirect(w, r, "/display", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/display", http.StatusSeeOther)
	}
}
func RemoveAdv(index int, save bool) {
	GetDataFromJson()
	var NewAdv []data.AdvStruct
	for _, adv := range Adventurer {
		if adv.Id != index {
			NewAdv = append(NewAdv, adv)
		}
	}
	Adventurer = NewAdv
	if save {
		SetDataToJson()
	}
}
