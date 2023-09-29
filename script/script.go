
package script

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/kmcsr/go-logger"
	"github.com/kmcsr/go-liter/script/console"
)

type Script struct {
	id      string
	version string
	path    string

	prog *goja.Program
	vm   *goja.Runtime
	loop *eventloop.EventLoop

	// exports *goja.Object
	doll    *goja.Object
	console *console.Console

	storage *MemoryStorage
	emitter *EventEmitter
}

func newScript(id string, version string, path string, prog *goja.Program, loop *eventloop.EventLoop)(*Script){
	return &Script{
		id: id,
		version: version,
		path: path,
		prog: prog,
		loop: loop,
	}
}

func (s *Script)Id()(string){
	return s.id
}

// TODO: Exports return a possible exports object from the script
// it's always nil for now, maybe it will be used later
func (s *Script)Exports()(*goja.Object){
	// return s.exports
	return nil
}

func (s *Script)Logger()(logger.Logger){
	return s.console.Logger()
}

func (s *Script)init(vm *goja.Runtime){
	if s.vm != nil {
		panic("liter: script: Script already initialized")
	}
	s.vm = vm
	s.storage = NewMemoryStorage(vm)
	s.emitter = NewEventEmitter(vm, s.loop)

	o := vm.NewObject()
	o.Set("ID", s.id)
	o.Set("VERSION", s.version)
	o.Set("storage", s.storage.Exports())
	s.emitter.ExportTo(o)
	s.doll = o
}

func (s *Script)On(name string, listener goja.Callable){
	s.emitter.OnAsync(name, listener)
}

func (s *Script)Off(name string, listener goja.Callable){
	s.emitter.OffAsync(name, listener)
}

func (s *Script)Emit(event *Event)(done <-chan bool){
	return s.emitter.EmitAsync(event)
}
