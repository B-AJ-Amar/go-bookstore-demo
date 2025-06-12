package main

import (
	"net/http"

	"github.com/B-AJ-Amar/go-bookstore-demo/bookstore/config"
	"github.com/B-AJ-Amar/go-bookstore-demo/bookstore/utils"
)

func main() {

	config.InitDB()
	utils.InitValidators()

	r := Router()

	http.ListenAndServe(":8000", r)
}