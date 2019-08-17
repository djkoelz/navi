package service

import (
	"github.com/djkoelz/navi/pkg/comm"
	"github.com/google/uuid"
)

type SlotFunc func(comm.Message)

type Slot struct {
	conn uuid.UUID // could make these uuids...
	f    SlotFunc
}

func NewSlot(f SlotFunc) *Slot {
	s := new(Slot)
	s.conn = uuid.New()
	s.f = f

	return s
}

func (this *Slot) Conn() uuid.UUID {
	return this.conn
}
