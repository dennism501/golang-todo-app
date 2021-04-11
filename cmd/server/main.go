package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/dennism501/golang-todo-app/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up our API")

	handler := transportHTTP.NewHandler()
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
