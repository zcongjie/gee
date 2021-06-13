package main

import (
	"net/http"

	"github.com/ijunyu/gee/engine"
)

func main() {
	r := engine.New()
	r.GET("/", func(c *engine.Context) {
		c.String(http.StatusOK, "welcome")
	})

	r.GET("/hello", func(c *engine.Context) {
		c.String(http.StatusOK, "hello, you're at %s\n", c.Path)
	})

	r.GET("/hello/:name", func(c *engine.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *engine.Context) {
		c.JSON(http.StatusOK, engine.H{"filepath": c.Param("filepath")})
	})

	r.Run(":8080")
}
