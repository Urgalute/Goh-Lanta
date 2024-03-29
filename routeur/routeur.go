package routeur

import (
	"Goh-Lanta/src"
	"fmt"
	"net/http"
	"os"
)

func InitServ() {

	http.HandleFunc("/", src.Home)
	http.HandleFunc("/add", src.FuncAdd)
	http.HandleFunc("/display", src.Display)
	http.HandleFunc("/adventurer", src.DisplayAdv)
	http.HandleFunc("/remove", src.Remove)
	http.HandleFunc("/updatepage", src.UpdatePage)
	http.HandleFunc("/update", src.Update)



	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:8080/) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)

}
