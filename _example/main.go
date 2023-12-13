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
	app.Get("/json/:name", Json)

	addr := fmt.Sprintf(":%d", PORT)

	log.Printf("Listening on %s...", addr)
	log.Fatal(app.Listen(addr))
}

func Index(ctx *miko.Context) error {
	return ctx.String("Hello from Index")
}

func Json(ctx *miko.Context) error {
	return ctx.Json(struct {
		Name string `json:"name"`
		Q    string `json:"q"`
	}{
		Name: ctx.RouteParam("name"),
		Q:    ctx.QueryString("q"),
	})
}
