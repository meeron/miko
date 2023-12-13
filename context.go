package miko

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	r *http.Request
	w http.ResponseWriter
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
