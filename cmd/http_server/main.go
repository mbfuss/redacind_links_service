package main

import (
	"1_increment_http_server/internal/bootstrap"
	"fmt"
)

func main() {
	init := bootstrap.App{}
	init.Initialize()
	err := init.StartServer()
	if err != nil {
		fmt.Println(err)
	}
}
