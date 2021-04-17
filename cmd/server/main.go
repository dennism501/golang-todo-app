package main

import (
	"fmt"
	"net/http"

	"github.com/dennism501/golang-todo-app/internal/database"
	"github.com/dennism501/golang-todo-app/internal/todo"
	transportHTTP "github.com/dennism501/golang-todo-app/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up our API")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	commentService := todo.NewService(db)
	handler := transportHTTP.NewHandler(commentService)
	handler.SetRouter()

	if err := http.ListenAndServe(":5000", handler.Router); err != nil {
		fmt.Println("Failed to set server :(")
		return err
	}
	return nil
}

func main() {

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting API server")
		fmt.Println(err)
	}

	http.ListenAndServe(":5000", nil)
}
