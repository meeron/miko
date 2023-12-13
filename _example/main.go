package main

import (
	"fmt"
	"github.com/meeron/miko"
	"log"
	"strconv"
)

var (
	PORT = 3000
)

func main() {
	app := miko.NewApp()
	app.Get("/", Index)
	app.Get("/json/:name", Json)
	app.Post("/post-json", PostBodyJson)
	app.Post("/post-form", PostBodyForm)

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

func PostBodyJson(ctx *miko.Context) error {
	body := struct {
		Name string `json:"name"`
	}{}

	if err := ctx.BindJson(&body); err != nil {
		return err
	}

	return ctx.Json(body)
}

func PostBodyForm(ctx *miko.Context) error {
	body := struct {
		Name string
		Age  int
	}{}

	age, err := strconv.Atoi(ctx.FormValue("age"))
	if err != nil {
		return err
	}

	body.Name = ctx.FormValue("name")
	body.Age = age

	return ctx.Json(body)
}
