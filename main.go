package main

import (
	"fmt"
	"module/src/core/domain/artist"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hi")
	artist := artist.Artist{
		ID:          uuid.New(),
		Name:        "Harry Styles",
		Age:         28,
		Nacionality: "American",
	}
	fmt.Println(artist)
}
