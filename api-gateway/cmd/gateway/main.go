package main

import "github.com/tanveerprottoy/starter-microservices/gateway/internal/app/gateway"

func main() {
	a := gateway.NewApp()
	a.Run()
}
