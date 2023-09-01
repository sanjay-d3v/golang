package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sanjay/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("server is geting started ...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listining at port 4000....")
}
