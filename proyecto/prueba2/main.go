package main

import "proyecto/prueba2/routes"

func main()  {
	r := routes.Init()

	r.Run(":3000")
}
