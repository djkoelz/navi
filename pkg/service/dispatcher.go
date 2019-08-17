package service

import (
	"github.com/djkoelz/navi/pkg/comm"
	"github.com/google/uuid"
)

type Dispatcher struct {
	slots map[comm.Description][]*Slot
}

func NewDispatcher() *Dispatcher {
	d := new(Dispatcher)
	d.slots = make(map[comm.Description][]*Slot, 0)
	return d
}

func (this *Dispatcher) Subscribe(desc comm.Description, f SlotFunc) uuid.UUID {
	// create a slot and append to list of slots at description
	slot := NewSlot(f)
	this.slots[desc] = append(this.slots[desc], slot)

	// retrn the slot connection id
	return slot.Conn()
}

func (this *Dispatcher) UnSubscribe(id uuid.UUID) {
	// we need to find a slot
	for desc, slots := range this.slots {
		for i, slot := range slots {
			if slot.Conn() == id {
				// remove without preserving order
				this.slots[desc] = remove(slots, i)
			}
		}
	}
}

func (this *Dispatcher) Post(m comm.Message) {
	slots, prs := this.slots[m.Desc()]
	if prs {
		for _, slot := range slots {
			go slot.f(m)
		}
	}
}

func remove(s []*Slot, i int) []*Slot {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
