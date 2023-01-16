package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/faisal-patel-456/mongoapi/Router"
)

func main() {
	fmt.Println("Mongo DB API")

	r := router.Router()
	fmt.Println("Server is getting started...")

	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening at port 4000 ...")
}
