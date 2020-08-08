package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Config the structure for the config JSON
type Config struct {
	Port string
	Host string
}

// MyResponse the response json strucutre which we will send from the api
type MyResponse struct {
	Status int
	Result interface{} // interface{} means ANYTHING is allowed
}

// Global variable (outside of main, can be used everywhere)
var config Config // Create config var with type Config we created up there

// HelloWorld example route of the API
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

// Home page of the api
func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "You're home! :D")
	// Send file over
	http.ServeFile(w, r, "home-page.html")
}

// PageStatus will return json of the page status
func PageStatus(w http.ResponseWriter, r *http.Request) {
	respToParse := MyResponse{Status: 200, Result: "great success!"}

	// Decode the json we want to send
	respToSend, err := json.Marshal(respToParse)

	if err != nil {
		panic("Couldn't read json :/")
	}
	// Set header so it can accept JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(respToSend)
}

func main() {
	// Read the config.json file
	bytes, err := ioutil.ReadFile("config.json")

	if err != nil {
		panic(err)
	}
	// Decode the JSON into the Config type we created
	json.Unmarshal(bytes, &config)

	// Create Routes
	http.HandleFunc("/helloworld", HelloWorld)
	http.HandleFunc("/", Home)
	http.HandleFunc("/status", PageStatus)

	fmt.Println("Api running on: " + config.Host + config.Port)

	// Launch the API
	http.ListenAndServe(config.Port, nil)
}
