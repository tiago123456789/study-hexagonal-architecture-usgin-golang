package main

import "github.com/tiago123456789/study-hexagonal-architecture-usgin-golang/adapters/web"

func main() {

	webServer := web.WebServer{}

	webServer.LoadRoutes()
	webServer.Start()
}
