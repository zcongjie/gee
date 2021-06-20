package main

import (
	"net/http"

	"github.com/ijunyu/gee/engine"
)

func main() {
	r := engine.Default()

	r.GET("/", func(c *engine.Context) {
		c.String(http.StatusOK, "welcome")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *engine.Context) {
			c.String(http.StatusOK, "hello, you're at %s\n", c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *engine.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		v2.GET("/assets/*filepath", func(c *engine.Context) {
			c.JSON(http.StatusOK, engine.H{"filepath": c.Param("filepath")})
		})
	}

	r.GET("/panic", func(c *engine.Context) {
		s := []string{"test"}
		c.String(http.StatusOK, s[100])
	})

	r.Run(":8080")
}
