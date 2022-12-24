package router

import "net/http"

type Router interface {
	Handle() http.Handler
}
