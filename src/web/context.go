package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
)

type Context struct {
	Params map[string]interface{} // 모든 자료형을 다 받기 위해 interface 사용

	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func (c *Context) RenderJson(v interface{}) {
	c.ResponseWriter.WriteHeader(http.StatusOK)
	c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(c.ResponseWriter).Encode(v); err != nil {
		c.RenderErr(http.StatusInternalServerError, err)
	}
}

func (c *Context) RenderErr(code int, err error) {
	if err != nil {
		if code > 0 {
			http.Error(c.ResponseWriter, http.StatusText(code), code)
		} else {
			defaultErr := http.StatusInternalServerError
			http.Error(c.ResponseWriter, http.StatusText(defaultErr), defaultErr)
		}
	}
}

var templates = map[string]*template.Template{}

func (c *Context) RenderTemplate(path string, v interface{}) {
	t, ok := templates[path]
	if !ok {
		t = template.Must(template.ParseFiles(filepath.Join(".", path)))
		templates[path] = t
	}

	t.Execute(c.ResponseWriter, v)
}

func (c *Context) Redirect(url string) {
	http.Redirect(c.ResponseWriter, c.Request, url, http.StatusMovedPermanently)
}
