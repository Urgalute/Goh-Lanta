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
	}
	temps.Temp.ExecuteTemplate(w, "add", nil)
}

// fonction pour ajouter un adv à notre tableau et potentiellement au json
func AddAdv(adventurer data.AdvStruct, save bool) {
	GetDataFromJson()
	Adventurer = append(Adventurer, adventurer)
	fmt.Println(adventurer)
	if save {
		SetDataToJson()
	}
}

// Fonction qui récup la longueur du tableau d'Adv +1 pour un nouvel article
func GetAdvId() int {
	id := len(Adventurer) + 1
	return id
}

func SetDataToJson() {
	data, err := json.Marshal(Adventurer)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	os.WriteFile("data/adv.json", data, 0644)
}

// fonction pour supprimer un adv de notre tableau et potentiellement du json
func RemoveAdv(index int, save bool) {
	GetDataFromJson()
	var NewArt []data.AdvStruct
	for _, art := range Adventurer {
		if art.Id != index {
			NewArt = append(NewArt, art)
		}
	}
	Adventurer = NewArt
	if save {
		SetDataToJson()
	}
}