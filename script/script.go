
package script

import (
	"io/fs"
	"encoding/json"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/kmcsr/go-logger"
	"github.com/kmcsr/go-liter/script/console"
)

type ScriptMeta struct {
	Id string `json:"name"`
	Version string `json:"version"`
	Description string `json:"description"`
	Dependencies DependMap `json:"dependencies"`
}

type Script struct {
	ScriptMeta

	loader *moduleLoader

	prog *goja.Program
	vm   *goja.Runtime
	loop *eventloop.EventLoop

	exports *goja.Object
	console *console.Console

	storage *MemoryStorage
	emitter *EventEmitter

	packet  fs.FS
	modules map[string]*goja.Object
}

func loadScriptMeta(packet fs.FS)(meta ScriptMeta, err error){
	fd, err := packet.Open("plugin.meta.json")
	if err != nil {
		return
	}
	defer fd.Close()
	if err = json.NewDecoder(fd).Decode(&meta); err != nil {
		return
	}
	return
}

func loadScript(packet fs.FS, meta ScriptMeta,
	extLoader extModuleLoader, loger logger.Logger,
	vm *goja.Runtime, loop *eventloop.EventLoop)(s *Script, err error){
	s = &Script{
		ScriptMeta: meta,
		vm: vm,
		loop: loop,
		emitter: NewEventEmitter(vm, loop),
	}

	doll := vm.NewObject()
	doll.Set("ID", s.Id)
	doll.Set("VERSION", s.Version)
	s.emitter.ExportTo(doll)
	s.console = console.NewConsole(vm, setPrefixLogger(loger, s.Id))
	s.loader = newModuleLoader(packet, vm, extLoader, []addonVar{
		{ name: "$", val: doll },
		{ name: "console", val: s.console.Exports() },
	})
	if s.exports, err = s.loader.load("index.js", "."); err != nil {
		return
	}
	return
}

// Exports return the module.exports from index.js
func (s *Script)Exports()(*goja.Object){
	return s.exports
}

func (s *Script)Logger()(logger.Logger){
	return s.console.Logger()
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
