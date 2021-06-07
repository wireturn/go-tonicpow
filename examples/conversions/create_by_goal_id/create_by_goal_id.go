package main

import (
	"log"
	"os"

	"github.com/tonicpow/go-tonicpow"
)

func main() {

	// Load the api client
	client, err := tonicpow.NewClient(
		tonicpow.WithAPIKey(os.Getenv("TONICPOW_API_KEY")),
		tonicpow.WithEnvironmentString(os.Getenv("TONICPOW_ENVIRONMENT")),
	)
	if err != nil {
		log.Fatalf("error in NewClient: %s", err.Error())
	}

	// Create conversion
	var conversion *tonicpow.Conversion
	conversion, err = client.CreateConversion(
		tonicpow.WithGoalID(13),
		tonicpow.WithTncpwSession("insert-your-visitor-tncpw-session-id"),
	)
	if err != nil {
		log.Fatalf("error in CreateConversion: %s", err.Error())
	}

	log.Printf("created conversion: %d", conversion.ID)
}
