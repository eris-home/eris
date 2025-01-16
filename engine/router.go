package engine

import "fmt"

type Router struct {
	Routes map[string]func(string)
}

func NewRouter() *Router {
	return &Router{
		Routes: make(map[string]func(string)),
	}
}

func (r *Router) RegisterRoute(command string, handler func(string)) {
	r.Routes[command] = handler
	fmt.Printf("Router: Registered command '%s'\n", command)
}

func (r *Router) GetHandler(command string) (func(string), bool) {
	handler, exists := r.Routes[command]
	return handler, exists
}