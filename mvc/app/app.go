package app

import (
	"net/http"

	"github.com/CodingMaven1/go-microservices/mvc/controller"
)

func StartApp() {
	http.HandleFunc("/users", controller.GetUser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
