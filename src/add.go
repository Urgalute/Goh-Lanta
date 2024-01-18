package src

import (
	"Goh-Lanta/data"
	"Goh-Lanta/temps"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Add(w http.ResponseWriter, r *http.Request) {
	temps.Temp.ExecuteTemplate(w, "add", nil)
}

// récupération des données du formulaire
func FuncAdd(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var Adventurer data.AdvStruct
		Adventurer.Name = r.FormValue("name")
		Adventurer.Content = r.FormValue("content")
		Adventurer.Team = r.FormValue("team")
		Adventurer.Preview = r.FormValue("preview")
		//Adventurer.Img = handler.Filename
		Adventurer.Id = GetAdvId()
		AddAdv(Adventurer, true)
		http.Redirect(w, r, "/display", http.StatusSeeOther)
	} else {
		temps.Temp.ExecuteTemplate(w, "add", nil)
	}
}

// fonction pour ajouter un adv à notre tableau et au json
func AddAdv(adventurer data.AdvStruct, save bool) {
	GetDataFromJson()
	Adventurer = append(Adventurer, adventurer)
	fmt.Println(adventurer)
	if save {
		SetDataToJson()
	}
}

func SetDataToJson() {
	data, err := json.Marshal(Adventurer)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	os.WriteFile("data/adv.json", data, 0644)
}
