package service

import (
	"github.com/djkoelz/navi/pkg/comm"
)

type SlotFunc func(comm.Message)

type Slot struct {
	conn int // could make these uuids...
	f    SlotFunc
}

func NewSlot(f SlotFunc) *Slot {
	s := new(Slot)
	s.conn = 0
	s.f = f

	return s
}
