package main

import (
	"github.com/shahen94/ddoser/internal"
)

func main() {
	coordinator := internal.NewCoordinator()

	finishChan := make(chan bool)

	go coordinator.Start(finishChan)

	<-finishChan
}
