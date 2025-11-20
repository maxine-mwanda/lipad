package main

import (
	"github.com/joho/godotenv"
	"lipad/controllers"
	"fmt"
)

var server = controllers.Server{}

func main() {
	fmt.Println("Hey there")
	_ = godotenv.Load()
	server.Init()
	fmt.Println("Debugging")
	server.Run()
	fmt.Println("Hey there")

}