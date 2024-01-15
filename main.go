package main

import (
	"Goh-Lanta/routeur"
	"Goh-Lanta/src"
	"Goh-Lanta/temps"
)

func main() {
	src.GetDataFromJson()
	temps.InitTemplate()
	routeur.InitServ()
}
