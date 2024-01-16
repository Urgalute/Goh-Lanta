package src

import (
	"Goh-Lanta/data"
	InitTemp "Goh-Lanta/temps"
	"net/http"
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
