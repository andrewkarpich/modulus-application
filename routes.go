package application

import (
	"context"
	"net/http"
)

type Action interface {
	NewRequestObject() any
	Handle(ctx context.Context, request any) (any, error)
}

type Routes struct {
	routes []RouteInfo
}

func NewRoutes() *Routes {
	return &Routes{routes: make([]RouteInfo, 0)}
}

func (r *Routes) Get(path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, RouteInfo{
		handler: handler,
		method:  http.MethodGet,
		path:    path,
	})
}

func (r *Routes) Post(path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, RouteInfo{
		handler: handler,
		method:  http.MethodPost,
		path:    path,
	})
}

func (r *Routes) Delete(path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, RouteInfo{
		handler: handler,
		method:  http.MethodDelete,
		path:    path,
	})
}

func (r *Routes) Put(path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, RouteInfo{
		handler: handler,
		method:  http.MethodPut,
		path:    path,
	})
}

func (r *Routes) Options(path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, RouteInfo{
		handler: handler,
		method:  http.MethodOptions,
		path:    path,
	})
}

func (r *Routes) AddFromRoutes(routes *Routes) {
	for name, info := range routes.routes {
		r.routes[name] = info
	}
}

func (r *Routes) GetRoutesInfo() []RouteInfo {
	result := make([]RouteInfo, 0, len(r.routes))
	for _, info := range r.routes {
		result = append(result, info)
	}

	return result
}
