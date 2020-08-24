package main

import (
	"fmt"
	"example-golang-postgres-app/router"
	"log"
	"net/http"
	"os"
)

func main() {
	r := router.Router()
	fmt.Sprintf("Starting server on the port %s...\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
}
