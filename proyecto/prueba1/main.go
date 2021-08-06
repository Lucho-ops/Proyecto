package main

import (
	"proyecto/prueba1/routes"
)

func main()  {
	r:= routes.Init()

	r.Run(":3000")
}

