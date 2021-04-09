package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Todo struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	CreatedAt string `json:"createdAt"`
}

var todoList []Todo

func init() {
	todoJSON := `[{"id":1,"message":"Create Todo list app with GoLang", "createdAt": "2015-12-12"}]`
	err := json.Unmarshal([]byte(todoJSON), &todoList)
	if err != nil {
		log.Fatal(err)
	}
}

func todoHandler(w http.ResponseWriter, r *http.Request) {

	todoListJson, err := json.Marshal(todoList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(todoListJson)
}

func main() {

	http.HandleFunc("/todos", todoHandler)
	http.ListenAndServe(":5000", nil)
}
