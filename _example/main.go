package main

import (
	"fmt"
	"github.com/meeron/miko"
	"log"
)

var (
	PORT = 3000
)

func main() {
	app := miko.NewApp()
	app.Get("/", Index)

	addr := fmt.Sprintf(":%d", PORT)

	log.Printf("Listening on %s...", addr)
	log.Fatal(app.Listen(addr))
}

func Index(ctx *miko.Context) error {
	return ctx.String("Hello from Index")
}
