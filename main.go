package main

import (
	"encoding/json"
	"fmt"
	"go-gorilla-crud/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/posts", getItem).Methods("GET")
	router.HandleFunc("/posts", addItems).Methods("POST")
	router.HandleFunc("/posts/{id}", getPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchItem).Methods("PATCH")
	router.HandleFunc("/posts/{id}", deleteItem).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

var data []models.Post = []models.Post{}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to interger"))

		return
	}

	if id >= len(data) {
		w.WriteHeader(404)
		w.Write([]byte("No data found with specified id"))
	}

	post := data[id]

	json.NewEncoder(w).Encode(post)
}

func addItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var newPost models.Post
	json.NewDecoder(r.Body).Decode(&newPost)

	data = append(data, newPost)

	json.NewEncoder(w).Encode(data)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	fmt.Println("Your details")

	json.NewEncoder(w).Encode(data)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to Integer"))
		return
	}

	//error checking

	if id >= len(data) {

		w.WriteHeader(404)
		w.Write([]byte("No data founded with  specified ID"))
		return

	}

	var updatedItem models.Post

	//updateItem := Post

	json.NewDecoder(r.Body).Decode(&updatedItem)

	data[id] = updatedItem

	json.NewEncoder(w).Encode(updatedItem)

}

func patchItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to Integer"))
		return
	}

	//error checking

	if id >= len(data) {

		w.WriteHeader(404)
		w.Write([]byte("No data founded with  specified ID"))
		return

	}
	// get the  current  value

	patchdata := data[id]

	json.NewDecoder(r.Body).Decode(&patchdata)

	data[id] = patchdata

	json.NewEncoder(w).Encode(patchdata)

}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to Integer"))
		return
	}

	//error checking

	if id >= len(data) {

		w.WriteHeader(404)
		w.Write([]byte("No data founded with  specified ID"))
		return

	}

	data = append(data[:id], data[id+1:]...)

	w.WriteHeader(200)
}
