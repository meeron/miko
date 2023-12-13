package miko

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type App struct {
	router *httprouter.Router
}

type Handler func(*Context) error

func NewApp() *App {
	return &App{
		router: httprouter.New(),
	}
}

func (app *App) Get(pattern string, h Handler) {
	app.router.GET(pattern, createHttpRouterHandler(h))
}

func (app *App) Listen(addr string) error {
	return http.ListenAndServe(addr, app.router)
}

func createHttpRouterHandler(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := &Context{
			w: w,
			r: r,
			p: p,
		}

		if err := h(ctx); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, err)
		}
	}
}
