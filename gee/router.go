package gee

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{make(map[string]HandlerFunc)}
}
