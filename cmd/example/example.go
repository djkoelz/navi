package main

import (
	"fmt"
	"github.com/djkoelz/navi/pkg/comm"
	"github.com/djkoelz/navi/pkg/service"
)

var signal = make(chan bool)

func echo(m comm.Message) {
	fmt.Println(m.Data())
	signal <- true
}

func main() {
	dispatcher := service.NewDispatcher()

	m := NewStringMessage("hey, listen...")

	id := dispatcher.Subscribe(StringMessage{}.Desc(), echo)

	fmt.Println(id)

	dispatcher.UnSubscribe(id)

	dispatcher.Post(m)

	<-signal
}
