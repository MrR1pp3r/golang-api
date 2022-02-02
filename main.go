package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing")
}

func handleresponse() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/{id}", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EndPoint Hit")
	json.NewEncoder(w).Encode(Articles)
}

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {
	Articles = []Article{
		{Id: "1", Title: "DEMO", Desc: "ARTICLE", Content: "LOL"},
		{Id: "2", Title: "DEMO1", Desc: "ARTICLE", Content: "LOL"},
	}

	handleresponse()
}
