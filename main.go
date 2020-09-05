package main

import(
	"flag"
	"github/hieutrtr/prime-service/handler"
	"github/hieutrtr/prime-service/router"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
	"github/hieutrtr/prime-service/store"
)

func main() {
	limitPrime := flag.Int64("limit-prime", 1000000, "limit of highest prime")
	primeCache := store.NewPrimeCache(uint32(*limitPrime))
	r := router.New()
	r.GET("/swagger/*", echoSwagger.WrapHandler)
	v1 := r.Group("/api/v1")
	h := handler.NewHandler(primeCache)
	h.RegisterV1(v1)
	r.Logger.Fatal(r.Start("0.0.0.0:8080"))
}
