package main

import (
	"fmt"
	"github.com/djkoelz/navi/pkg/comm"
	"github.com/djkoelz/navi/pkg/control"
	"github.com/djkoelz/navi/pkg/service"
	"time"
)

func echo(m comm.Message) {
	fmt.Println(m.Data())
}

func main() {
	dispatcher := service.NewDispatcher()

	m := NewStringMessage("hey, listen...")

	waiter := control.NewWaiter(StringMessage{}.Desc())

	dispatcher.Subscribe(StringMessage{}.Desc(), echo)

	go func() {
		if !waiter.Wait(dispatcher, 2*time.Second) {
			fmt.Println("Timedout...")
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		dispatcher.Post(m)
	}()

	for {
	}
}
