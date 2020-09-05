package main

import(
	"flag"
	"github.com/hieutrtr/prime-service/handler"
	"github.com/hieutrtr/prime-service/router"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
	"github.com/hieutrtr/prime-service/store"
)

func main() {
	limitPrime := flag.Int64("limit-prime", 1000000, "limit of highest prime")
	JWTSecret := flag.String("jwt-secret", "SECRET", "secret to generate JWT")
	primeCache := store.NewPrimeCache(uint32(*limitPrime))
	r := router.New()
	r.GET("/swagger/*", echoSwagger.WrapHandler)
	v1 := r.Group("/api/v1")
	h := handler.NewHandler(primeCache)
	h.RegisterV1(v1, *JWTSecret)
	r.Logger.Fatal(r.Start("0.0.0.0:8080"))
}
