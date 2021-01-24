package app

import (
	"fmt"
	"net/http"

	"github.com/nakiscia/go-microservice-example/controllers"
)

func StartApplication() {

	handleRequests()

	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}

	fmt.Println("Application started on :8080")
}

func handleRequests() {
	// Handler for /user
	http.HandleFunc("/user", controllers.GetUser)
}
