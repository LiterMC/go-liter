
package script

import (
	"sync"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/kmcsr/go-logger"
	"github.com/kmcsr/go-liter/script/console"
)

type JSObject interface {
	// we have a different signature than goja.Value has `ToObject(*goja.Runtime)(*goja.Object)`
	AsObject(*goja.Runtime)(*goja.Object)
}

type Script struct {
	id      string
	version string
	path    string

	prog *goja.Program
	loop *eventloop.EventLoop

	// exports *goja.Object
	doll    *goja.Object
	console *console.Console

	storage *MemoryStorage

	listenerMux sync.Mutex
	listeners   map[string][]goja.Callable
}

func newScript(id string, version string, path string, prog *goja.Program, loop *eventloop.EventLoop)(*Script){
	return &Script{
		id: id,
		version: version,
		path: path,
		prog: prog,
		loop: loop,
		listeners: make(map[string][]goja.Callable),
	}
}

func (s *Script)Id()(string){
	return s.id
}

// Exports return a possible exports object which returned by the script
// Now it's always nil
func (s *Script)Exports()(*goja.Object){
	// return s.exports
	return nil
}

func (s *Script)Logger()(logger.Logger){
	return s.console.Logger()
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

func (s *Script)Emit(event string, args ...any){
	if listeners, ok := s.listeners[event]; ok {
		if len(listeners) == 0 {
			return
		}
		done := make(chan struct{}, 0)
		s.loop.RunOnLoop(func(vm *goja.Runtime){
			defer close(done)
			jsArgs := wrapToJSValues(vm, args...)
			for _, l := range listeners {
				l(nil, jsArgs...)
			}
		})
		<-done
	}
}

func (s *Script)EmitAsync(event string, args ...any)(done <-chan struct{}){
	exit := make(chan struct{}, 0)
	if listeners, ok := s.listeners[event]; ok {
		if len(listeners) == 0 {
			close(exit)
			return
		}
		s.loop.RunOnLoop(func(vm *goja.Runtime){
			defer close(exit)
			jsArgs := wrapToJSValues(vm, args...)
			for _, l := range listeners {
				l(nil, jsArgs...)
			}
		})
	}else{
		close(exit)
	}
	return exit
}

func wrapToJSValues(vm *goja.Runtime, args ...any)(res []goja.Value){
	res = make([]goja.Value, len(args))
	for i, v := range args {
		if v, ok := v.(JSObject); ok {
			res[i] = v.AsObject(vm)
		}else{
			res[i] = vm.ToValue(v)
		}
	}
	return
}
