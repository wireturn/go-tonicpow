package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/tonicpow/go-tonicpow"
)

var (
	// TonicPowAPI is the client we will load on start-up
	TonicPowAPI *tonicpow.Client
)

// Load the TonicPow API Client once when the application loads
func init() {

	//
	// Get the API key (from env or your own config)
	//
	apiKey := os.Getenv("TONICPOW_API_KEY")
	if len(apiKey) == 0 {
		log.Fatalf("api key is required: %s", "TONICPOW_API_KEY")
	}

	//
	// Load the api client (creates a new session)
	// You can also set the environment or client options
	//
	var err error
	TonicPowAPI, err = tonicpow.NewClient(apiKey, tonicpow.LocalEnvironment, nil)
	if err != nil {
		log.Fatalf("error in NewClient: %s", err.Error())
	}
}

func main() {
	//
	// Example for ending the api session for the application
	// This is not needed, sessions will automatically expire
	//
	defer func() {
		_ = TonicPowAPI.EndSession("")
	}()

	// Example vars
	var err error
	var userSessionToken string
	testPassword := "ExamplePassForNow0!"

	//
	// Example: Prolong a session
	//
	if err = TonicPowAPI.ProlongSession(""); err != nil {
		log.Fatalf("ProlongSession: %s", err.Error())
	} else {
		log.Println("session created and prolonged...")
	}

	//
	// Example: Create a user
	//
	user := &tonicpow.User{
		Email:    fmt.Sprintf("Testing%d@TonicPow.com", rand.Intn(100000)),
		Password: testPassword,
	}
	if user, err = TonicPowAPI.CreateUser(user); err != nil {
		log.Fatalf("create user failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	} else {
		log.Printf("user %d created", user.ID)
	}

	//
	// Example: Get a user (id)
	//
	if user, err = TonicPowAPI.GetUser(user.ID, ""); err != nil {
		log.Fatalf("get user failed - api error: %s", TonicPowAPI.LastRequest.Error.Message)
	} else {
		log.Printf("got user by id %d", user.ID)
	}

	//
	// Example: Get a user (email)
	//
	if user, err = TonicPowAPI.GetUser(0, user.Email); err != nil {
		log.Fatalf("get user failed - api error: %s", TonicPowAPI.LastRequest.Error.Message)
	} else {
		log.Printf("got user by email %s", user.Email)
	}

	//
	// Example: Update a user
	//
	user.FirstName = "Austin"
	if user, err = TonicPowAPI.UpdateUser(user, ""); err != nil {
		log.Fatalf("update user failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	} else {
		log.Printf("user %d updated - first_name: %s", user.ID, user.FirstName)
	}

	//
	// Example: Login for a user
	//
	user.Password = testPassword
	userSessionToken, err = TonicPowAPI.LoginUser(user)
	if err != nil {
		log.Fatalf("user login failed - api error: %s data: %s", TonicPowAPI.LastRequest.Error.Message, TonicPowAPI.LastRequest.Error.Data)
	} else {
		log.Printf("user login: %s token: %s", user.Email, userSessionToken)
	}
}
