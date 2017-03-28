package main

import (
	"fmt"
	"html"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	fmt.Println("Launching Demo web Server")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The beer of the day is:  %q\n", html.EscapeString(generateBeer()))
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}

func generateBeer() string {

	beers := []string{
		"Krombacher",
		"Bitburger",
		"Warsteiner",
		"Becks",
		"Veltins",
		"Hasseröder",
		"Paulaner Weissbier",
		"Radeberger",
		"König Pilsener",
		"Erdinger Weissbier",
		"Holsten Pilsener",
		"Jever",
		"Franziskaner Weissbier",
		"Diebels Alt",
		"Oettinger",
		"Wernesgrüner",
		"Berliner Pilsener",
		"Lübzer Pils",
		"Köstritzer Schwarzbier",
		"Flensburger Pilsener",
	}

	beerIndex := rand.Intn(len(beers))

	return beers[beerIndex]
}
