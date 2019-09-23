package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Errorf("", err)
		}
	}()

	go gogogo()
	var t chan struct{}
	<-t
}

func gogogo() {
	panic(nil)
}
