package routes

import "net/http"


type Middleware func(http.HandlerFunc) http.HandlerFunc

type Router struct {
  mux *http.ServeMux
  middlewares []Middleware
}

func NewRouter() *Router {
  return &Router{
    mux: http.NewServeMux(),
  }
}

func (r *Router) Use(middlewares Middleware) {
  r.middlewares = append(r.middlewares, middlewares)
} 

func (r *Router) HandleFunc(pattern string, handler http.HandlerFunc) {
  finalHandler := r.applyMiddleware(handler)
  r.mux.HandleFunc(pattern, finalHandler)
}

func (r *Router) applyMiddleware(handler http.HandlerFunc) http.HandlerFunc {
  for i := len(r.middlewares) - 1; i >= 0; i-- {
    handler = r.middlewares[i](handler)
  }

  return handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  r.mux.ServeHTTP(w, req)
}