package server

import (
	"fmt"
	"net/http"
)

type api struct{}

func (a *api) ServeHTTP(http.ResponseWriter, *http.Request) {
	fmt.Println("hello, I'm from Go")
}

func Run() error {
	http.Handle("/api", &api{})
	return http.ListenAndServe(":8080", nil)
}
