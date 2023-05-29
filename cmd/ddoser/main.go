package main

import (
	"github.com/shahen94/ddoser/internal"
)

func checkError() {
	if rec := recover(); rec != nil {
		return
	}
}

func main() {
	defer checkError()
	coordinator := internal.NewCoordinator()

	finishChan := make(chan bool)

	go coordinator.Start(finishChan)

	<-finishChan
}
