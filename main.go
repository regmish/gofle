package main

import (
	"fmt"
	"gofle/internal/http"
	"os"
)

const API_URL = "https://api.regmi.de/v1"

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Owner string `json:"ownerId"`
}

func main() {
	var email = "shankarregmi@gmail.com"
	var password = "password"
	tokenBody := &struct {
		Token string `json:"token"`
	}{}

	http.Post(API_URL+"/auth/login", []byte(fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password)), &tokenBody)
	os.Setenv("AUTHORIZATION_TOKEN", tokenBody.Token)

	// fmt.Println(tokenBody.Token)

	// Fetch notes

	var res []Note
	http.Get(API_URL+"/notes", &res)
	for _, note := range res {
		fmt.Println(note.Title)
		// 	// fmt.Println(note.Body)
		// fmt.Println("************************************************************************")
	}

	// Create note

	// noteBody, _ := json.Marshal(&Note{
	// 	Title: time.Now().String(),
	// 	Body:  "This note is created from Gofle, that works",
	// })

	// response := &struct {
	// 	Id string `json:"_id"`
	// 	Note
	// }{}

	// generator.Generate(&generator.GeneratorArgs{
	// 	Url:      API_URL + "/notes",
	// 	Body:     noteBody,
	// 	Response: &response,
	// })

	// fmt.Println(response.Id)
	// fmt.Println(response.Title)
	// fmt.Println(response.Body)
}
