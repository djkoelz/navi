package main

import (
	"fmt"
	"github.com/djkoelz/navi/pkg/comm"
	"github.com/djkoelz/navi/pkg/service"
	"reflect"
)

type StringMessage struct {
	data string
}

func NewStringMessage(d string) *StringMessage {
	m := new(StringMessage)
	m.data = d

	return m
}

func (this StringMessage) Desc() comm.Description {
	return reflect.TypeOf(this)
}

func (m StringMessage) Data() interface{} {
	return m.data
}

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
