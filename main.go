package main

import (
	"flag"
	"log"

	"./api"
)

func main() {
	bindAddr := flag.String("bind_addr", ":8080", "Set bind address")
	flag.Parse()
	a := api.New(":9111")
	log.Fatal(a.Start())
}
