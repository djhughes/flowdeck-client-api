package main

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/djhughes/flowdeck-client-api/operations"
)

func main() {
	router := http.NewServeMux()
	api := humago.New(router, huma.DefaultConfig("Runflow API", "1.0.0"))

	runbooksHandler := &operations.RunbooksHandler{}
	huma.AutoRegister(api, runbooksHandler)

	http.ListenAndServe(":8080", router)
}
