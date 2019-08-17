package service

import (
	"github.com/djkoelz/navi/pkg/comm"
)

type Dispatcher struct {
	slots map[comm.Description][]*Slot
}

func NewDispatcher() *Dispatcher {
	d := new(Dispatcher)
	d.slots = make(map[comm.Description][]*Slot, 0)
	return d
}

func (this *Dispatcher) Subscribe(desc comm.Description, f SlotFunc) {
	this.slots[desc] = append(this.slots[desc], NewSlot(f))
}

func (this *Dispatcher) Post(m comm.Message) {
	slots, prs := this.slots[m.Desc()]
	if prs {
		for _, slot := range slots {
			go slot.f(m)
		}
	}
}
