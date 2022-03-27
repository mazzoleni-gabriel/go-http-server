package main

import (
	"github.com/mazzoleni-gabriel/go-http-server/user/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/user", handler.Handler)

	http.ListenAndServe(":8090", nil)
}
