package main

import (
	"basic-api-docker/infra/handler"
)

func main() {
	router := handler.NewRouter()
	router.Logger.Fatal(router.Start(":8080"))
}
