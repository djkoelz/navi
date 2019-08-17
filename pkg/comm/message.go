package comm

import (
	"reflect"
)

type Description reflect.Type

type Message interface {
	Desc() Description // gets the type descriptor for this event
	Data() interface{} // gets the data packaged with the event
}
