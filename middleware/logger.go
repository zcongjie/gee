package middleware

import (
	"log"
	"time"

	"github.com/ijunyu/gee/engine"
)

func Logger() engine.HandlerFunc {
	return func(c *engine.Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
