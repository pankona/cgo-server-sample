package server

import (
	"fmt"
	"net/http"
)

type api struct {
	callback func() string
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received. hello, I'm from Go!")

	w.Write([]byte("hello, I'm from Go!\n"))

	// write response with buffer from callback
	buf := a.callback()
	w.Write([]byte(buf))
}

func Run(f func() string) error {
	http.Handle("/api", &api{
		callback: f,
	})
	return http.ListenAndServe(":8080", nil)
}
