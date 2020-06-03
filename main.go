package main

import (
	"log"

	"./api"
)

func main() {
	a := api.New(":9111")
	log.Fatal(a.Start())
}
