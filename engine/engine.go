package engine

import "net/http"

type HandlerFunc func(*Context)

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {}
