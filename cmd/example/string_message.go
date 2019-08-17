package main

import (
	"github.com/djkoelz/navi/pkg/comm"
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
