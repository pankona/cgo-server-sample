package main

import (
	"fmt"

	"github.com/pankona/cgo-server-sample/server"
)

func main() {
	fmt.Println(server.Run())
}

//export Run
func Run() error {
	return server.Run()
}
