package control

import (
	"github.com/djkoelz/navi/pkg/comm"
	"github.com/djkoelz/navi/pkg/service"
	"time"
)

// should wait for an message to be posted
// add a timeout feature, if the waiter timesout, we have failed the wait
// otherwise success
type Waiter struct {
	desc   comm.Description
	signal chan bool
}

func NewWaiter(d comm.Description) *Waiter {
	w := new(Waiter)
	w.desc = d
	w.signal = make(chan bool)

	return w
}

func (this *Waiter) callback(comm.Message) {
	this.signal <- true
}

func (this *Waiter) Wait(d *service.Dispatcher, timeout time.Duration) bool {
	d.Subscribe(this.desc, this.callback)

	// wait with timeout until fails
	select {
	case <-this.signal:
		return true
	case <-time.After(timeout):
		return false
	}

	return <-this.signal
}
