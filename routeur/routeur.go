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
	http.HandleFunc("/adventurer", src.FuncAdd)



	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:8080/) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)

}
