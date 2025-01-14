package http

import "net/http"

func registerRouter() *http.ServeMux {
	r := http.NewServeMux()
	return r
}
