package main

import (
	"fmt"
	"github.com/djkoelz/navi/pkg/comm"
	"github.com/djkoelz/navi/pkg/service"
)

func echo(m comm.Message) {
	fmt.Println(m.Data())
}

func main() {
	dispatcher := service.NewDispatcher()

	m := NewStringMessage("hey, listen...")

	dispatcher.Subscribe(StringMessage{}.Desc(), echo)
	dispatcher.Post(m)

	for {
	}
}
