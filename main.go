package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request)  {
	articles := Articles{
		Article{Title: "test Title", Desc: "Test Description", Content: "Test content"},
	}
	fmt.Println("Endpoint Hit: All articles hit ")
	json.NewEncoder(w).Encode(articles)
	
}

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests(){

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/article/", allArticles).Methods("GET")
	myRouter.HandleFunc("/article/", testPostArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func testPostArticle(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Test POST Endpoint Hit")
}

func main(){
	handleRequests()
}
