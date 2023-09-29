
package main

import (
	"github.com/dop251/goja"
	"github.com/kmcsr/go-liter"
	"github.com/kmcsr/go-liter/script"
)

type Event struct {
	Name string
	Client *liter.Conn
	Server *liter.Conn
	Data map[string]any
}

var _ script.JSObject = (*Event)(nil)

func (e *Event)AsObject(vm *goja.Runtime)(o *goja.Object){
	o = vm.NewObject()
	o.Set("name", e.Name)
	o.Set("client", e.Client)
	o.Set("server", e.Server)
	o.Set("data", e.Data)
	for k, v := range e.Data {
		switch k {
		case "name", "client", "server", "data": // reserve keys
			continue
		}
		o.Set(k, v)
	}
	return
}
