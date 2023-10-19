
package script

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
)

type Exportsable interface {
	Exports()(*goja.Object)
}

type Event struct {
	Name string
	Data map[string]any
	Cancelable bool

	cancelled bool
}

// NewEvent will create a cancelable event
func NewEvent(name string, data map[string]any)(*Event){
	return &Event{
		Name: name,
		Data: data,
		Cancelable: true,
	}
}

func (e *Event)toValue(vm *goja.Runtime, loop *eventloop.EventLoop)(goja.Value){
	if e == nil {
		return goja.Undefined()
	}
	o := vm.NewObject()
	o.Set("name", e.Name)
	o.Set("cancel", e.Cancel)
	o.DefineAccessorProperty("cancelable", vm.ToValue(func(goja.FunctionCall)(goja.Value){
		return vm.ToValue(e.Cancelable)
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	o.DefineAccessorProperty("cancelled", vm.ToValue(func(goja.FunctionCall)(goja.Value){
		return vm.ToValue(e.Cancelled())
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	if e.Data == nil {
		o.Set("data", vm.NewObject()) // a empty object
	}else{
		o.Set("data", e.Data)
		for k, v := range e.Data {
			switch k { // ignore reserved keys
			case "name", "data", "cancel", "cancelable", "cancelled":
				continue
			}
			switch v := v.(type) {
			case Exportsable:
				o.Set(k, v.Exports())
			case func(vm *goja.Runtime)(goja.Value):
				o.Set(k, v(vm))
			case func(vm *goja.Runtime, loop *eventloop.EventLoop)(goja.Value):
				o.Set(k, v(vm, loop))
			default:
				o.Set(k, v)
			}
		}
	}
	return o
}

func (e *Event)Cancel(){
	if e.Cancelable {
		e.cancelled = true
	}
}

func (e *Event)Cancelled()(bool){
	return e.cancelled
}
