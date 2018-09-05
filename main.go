package main

import (
	"GO_Admin/router"
)

func main() {
	r := route.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
