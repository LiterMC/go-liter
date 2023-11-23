package script

import (
	"reflect"
	"sync"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
)

type EventEmitter struct {
	vm   *goja.Runtime
	loop *eventloop.EventLoop

	listenerMux sync.RWMutex
	listeners   map[string][]goja.Callable
}

func NewEventEmitter(vm *goja.Runtime, loop *eventloop.EventLoop) (e *EventEmitter) {
	return &EventEmitter{
		vm:        vm,
		loop:      loop,
		listeners: make(map[string][]goja.Callable),
	}
}

func (e *EventEmitter) ExportTo(o *goja.Object) {
	o.Set("on", e.on)
	o.Set("addListener", e.on)
	o.Set("off", e.off)
	o.Set("removeListener", e.off)
	o.Set("emit", func(call goja.FunctionCall, vm *goja.Runtime) goja.Value {
		event := &Event{
			Name:       call.Arguments[0].String(),
			Cancelable: false,
		}
		if len(call.Arguments) > 1 {
			arg1 := call.Argument(1).Export()
			var ok bool
			if event.Data, ok = arg1.(map[string]any); ok {
				if len(call.Arguments) > 2 {
					event.Cancelable = call.Argument(2).ToBoolean()
				}
			} else if cancelable, ok := arg1.(bool); ok {
				event.Cancelable = cancelable
			}
		}
		return vm.ToValue(e.emit(event))
	})
	o.Set("removeAllListeners", e.removeAllListeners)
	o.Set("eventNames", e.eventNames)
	o.Set("listeners", e.getListeners)
}

func (e *EventEmitter) on(name string, listener goja.Callable) *EventEmitter {
	e.listeners[name] = append(e.listeners[name], listener)
	return e
}

func (e *EventEmitter) OnAsync(name string, listener goja.Callable) {
	e.loop.RunOnLoop(func(vm *goja.Runtime) {
		e.on(name, listener)
	})
}

func (e *EventEmitter) off(name string, listener goja.Callable) {
	if listeners, ok := e.listeners[name]; ok {
		hp := reflect.ValueOf(listener).Pointer()
		for i := len(listeners) - 1; i >= 0; i-- {
			if reflect.ValueOf(listeners[i]).Pointer() == hp {
				copy(listeners[i:], listeners[i+1:])
				e.listeners[name] = listeners[:len(listeners)-1]
				break
			}
		}
	}
}

func (e *EventEmitter) OffAsync(name string, listener goja.Callable) {
	e.loop.RunOnLoop(func(vm *goja.Runtime) {
		e.off(name, listener)
	})
}

func (e *EventEmitter) removeAllListeners(name string) {
	delete(e.listeners, name)
}

func (e *EventEmitter) emit(event *Event) (cancelled bool) {
	if event.Cancelled() {
		panic(e.vm.ToValue("Error: Trying to emit a cancelled event"))
	}
	if listeners, ok := e.listeners[event.Name]; ok {
		if len(listeners) == 0 {
			return false
		}
		arg := event.toValue(e.vm, e.loop)
		for _, l := range listeners {
			if _, err := l(nil, arg); err != nil {
				panic(e.vm.ToValue(err))
			}
			if event.Cancelled() {
				return true
			}
		}
	}
	return false
}

func (e *EventEmitter) EmitAsync(event *Event) (done <-chan bool) {
	exit := make(chan bool, 1)
	e.loop.RunOnLoop(func(vm *goja.Runtime) {
		defer close(exit)
		exit <- e.emit(event)
	})
	return exit
}

func (e *EventEmitter) eventNames() (names []string) {
	names = make([]string, 0, len(e.listeners))
	for n, _ := range e.listeners {
		names = append(names, n)
	}
	return
}

func (e *EventEmitter) getListeners(name string) (handlers []goja.Callable) {
	listeners := e.listeners[name]
	handlers = make([]goja.Callable, 0, len(listeners))
	for _, c := range listeners {
		handlers = append(handlers, c)
	}
	return
}
