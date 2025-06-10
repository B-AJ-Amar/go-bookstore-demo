package main

import (
	"net/http"

	"github.com/B-AJ-Amar/go-bookstore-demo/bookstore/config"
)

func main() {

	config.InitDB()

	r := Router()

	http.ListenAndServe(":8000", r)
}