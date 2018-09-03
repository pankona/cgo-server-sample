package main

import (
	"fmt"

	"github.com/pankona/cgo-server-sample/server"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(server.Run())
}
