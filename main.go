package main

import (
	"github.com/sureshtamrakar/socials/routes"
)

func main() {
	r := routes.AddRoutes()
	r.Run()

}
