package main

import (
	"crud-movies/routes"
)

func main() {
	//Route init
	e := routes.InitRoute()
	
	e.Start(":4000")
}