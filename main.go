package main

import(
	"flag"
	"github.com/hieutrtr/prime-service/handler"
	"github.com/hieutrtr/prime-service/router"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
	_ "github.com/hieutrtr/prime-service/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/hieutrtr/prime-service/store"
)

// @title Prime API
// @version 1.0
// @description A service for finding nearest prime less than a given number N.

// @contact.name Hieu TRAN
// @contact.email hieutrantrung.it@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
