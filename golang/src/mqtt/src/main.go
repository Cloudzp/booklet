package main

import (
	"mqtt/server"
)

func main() {

	svr, err := server.NewServer("localhost:8888")
	if err != nil {
		return
	}
	svr.Run()
}
