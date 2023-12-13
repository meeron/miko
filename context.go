package miko

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Context struct {
	r *http.Request
	w http.ResponseWriter
	p httprouter.Params
}

func (c *Context) RouteParam(name string) string {
	return c.p.ByName(name)
}

func (c *Context) QueryString(name string) string {
	return c.r.URL.Query().Get(name)
}

func (c *Context) String(text string) error {
	c.w.Header().Add("Content-Type", "text/plain")
	_, err := fmt.Fprint(c.w, text)
	return err
}

func (c *Context) StatusJson(data any, statusCode int) error {
	c.w.Header().Add("Content-Type", "application/json")
	c.w.WriteHeader(statusCode)

	enc := json.NewEncoder(c.w)
	return enc.Encode(data)
}

func (c *Context) Json(data any) error {
	return c.StatusJson(data, http.StatusOK)
}

func (c *Context) BindJson(v any) error {
	if c.r.Header.Get("Content-Type") != "application/json" {
		return ErrUnsupportedMediaType
	}

	enc := json.NewDecoder(c.r.Body)
	return enc.Decode(v)
}

func (c *Context) FormValue(name string) string {
	return c.r.FormValue(name)
}
