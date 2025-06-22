package main

import (
	"net/http"
	"context"
	"log"

	"github.com/B-AJ-Amar/go-bookstore-demo/bookstore/config"
	"github.com/B-AJ-Amar/go-bookstore-demo/bookstore/utils"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"github.com/B-AJ-Amar/go-bookstore-demo/otel"
)

func main() {

	ctx := context.Background()
	shutdown := otel.InitTracer("bookstore-service")
	defer shutdown(ctx)

	config.InitDB()
	utils.InitValidators()
	

	r := Router()

	

	http.ListenAndServe(":8000", r)
}