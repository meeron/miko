package miko

import (
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
