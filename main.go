package main

import (
	"sync"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var requestCodes = &requestCodeMap{
	m: map[string]string{},
}

type requestCodeMap struct {
	mu sync.RWMutex
	m  map[string]string
}

func (r *requestCodeMap) Get(k string) string {
	r.mu.Lock()
	defer r.mu.Unlock()
	if v, ok := r.m[k]; ok {
		return v
	}

	return ""
}

func (r *requestCodeMap) Set(k, v string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.m[k] = v
}

func main() {
	e := echo.New()

	e.GET("/", top)
	e.GET("/pocket/auth", auth)
	e.GET("/pocket/redirect", redirect)
	e.GET("/pocket/token", token)
	e.GET("/pocket/list", list)
	e.GET("/pocket", redirectToAuth)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":1323"))
}
