
package script

import (
	"sync"

	"github.com/dop251/goja"
	"github.com/kmcsr/go-liter/script/console"
)

type Script struct {
	id      string
	version string
	path string
	prog *goja.Program
	doll *goja.Object
	console *console.Console

	storage *MemoryStorage

	listenerMux sync.Mutex
	listeners   map[string][]goja.Callable
}

func newScript(id string, version string, path string, prog *goja.Program)(*Script){
	return &Script{
		id: id,
		version: version,
		path: path,
		prog: prog,
		listeners: make(map[string][]goja.Callable),
	}
}

func (s *Script)Id()(string){
	return s.id
}

func (s *Script)init(vm *goja.Runtime){
	if s.doll != nil {
		panic("liter: script: dollar object already initialized")
	}
	s.storage = NewMemoryStorage(vm)

	o := vm.NewObject()
	o.Set("ID", s.id)
	o.Set("VERSION", s.version)
	o.Set("on", s.js_on)
	o.Set("storage", s.storage.Exports())
	s.doll = o
}

func (s *Script)js_on(call goja.FunctionCall, vm *goja.Runtime)(goja.Value){
	event0, ok := call.Argument(0).(goja.String)
	if !ok {
		panic(vm.ToValue("arg#0 is not a string of event id"))
	}
	event := event0.String()
	cb, ok := goja.AssertFunction(call.Argument(1))
	if !ok {
		panic(vm.ToValue("arg#1 is not a function"))
	}
	s.listenerMux.Lock()
	defer s.listenerMux.Unlock()
	s.listeners[event] = append(s.listeners[event], cb)
	return nil
}
