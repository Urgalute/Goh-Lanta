package src

import (
	"Goh-Lanta/data"
	InitTemp "Goh-Lanta/temps"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var Adventurer []data.AdvStruct

// mainfunc pour la page home
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/flavicon.ico" {
		InitTemp.Temp.ExecuteTemplate(w, "error404", nil)
		return
	}
	InitTemp.Temp.ExecuteTemplate(w, "home", Adventurer /*RandomArticle(4)*/)

}

func GetDataFromJson() {
	data, err := os.ReadFile("data/adv.json") //ouverture et lecture du json
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	json.Unmarshal(data, &Adventurer) //passage en json vers la struct
}
