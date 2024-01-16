package src

import (
	"Goh-Lanta/data"
	"encoding/json"
	"fmt"
	"os"
)

func GetDataFromJson() {
	data, err := os.ReadFile("data/adv.json") //ouverture et lecture du json
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	json.Unmarshal(data, &Adventurer) //passage en json vers la struct
}

// Fonction qui récup la longueur du tableau d'Adv +1 pour un nouvel article
func GetAdvId() int {
	id := len(Adventurer) + 1
	return id
}


// fonction pour récupèrer tous les articles toutes catégories confondues
func GetAllAdv() []data.AdvStruct {
	GetDataFromJson()
	return Adventurer
}

func AdvExist(id int) bool {
	GetDataFromJson()
	for _, adv := range Adventurer {
		if adv.Id == id {
			return true
		}
	}
	return false
}