package main

import (
	"log"
	// "thirdPartLogin/lib"
	"thirdPartLogin/router"
)

func main() {
	// serverConfig := lib.LoadServerConfig()

	r := router.SetupRoute()

	r.LoadHTMLGlob("view/*")

	if err := r.Run(":80"); err != nil {
		log.Fatal("server start error...")
	}
}
