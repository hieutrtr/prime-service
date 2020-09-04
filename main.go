package main

import(
	"prime-service/router"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

func main() {
	r := router.New()
	r.GET("/swagger/*", echoSwagger.WrapHandler)
}
