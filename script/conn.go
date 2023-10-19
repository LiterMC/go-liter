
package script

import (
	"sync/atomic"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/kmcsr/go-liter"
)

type WrappedConn struct {
	*liter.Conn
	vm *goja.Runtime
	loop *eventloop.EventLoop
	closed atomic.Bool

	emitter *EventEmitter
	exports *goja.Object
}

var _ Exportsable = (*WrappedConn)(nil)

func WrapConn(conn *liter.Conn, vm *goja.Runtime, loop *eventloop.EventLoop)(c *WrappedConn){
	c = &WrappedConn{
		Conn: conn,
		vm: vm,
		loop: loop,
		emitter: NewEventEmitter(vm, loop),
	}
	o := vm.NewObject()
	o.DefineAccessorProperty("protocol", vm.ToValue(func(goja.FunctionCall)(goja.Value){
		return vm.ToValue(conn.Protocol())
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	o.DefineAccessorProperty("localAddr", vm.ToValue(func(goja.FunctionCall)(goja.Value){
		return vm.ToValue(conn.LocalAddr().String())
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	o.DefineAccessorProperty("remoteAddr", vm.ToValue(func(goja.FunctionCall)(goja.Value){
		return vm.ToValue(conn.RemoteAddr().String())
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	o.Set("close", vm.ToValue(func(goja.FunctionCall)(goja.Value){
		c.closed.Store(true)
		conn.Close()
		return goja.Undefined()
	}))
	o.Set("newPacket", func(call goja.FunctionCall)(goja.Value){
		if len(call.Arguments) == 0 {
			panic(vm.NewTypeError("arg#0 is not a integer of packet id"))
		}
		id := (liter.VarInt)(call.Argument(0).ToInteger())
		return WrapPacketBuilder(liter.NewPacket(conn.Protocol(), id), conn, vm)
	})
	c.emitter.ExportTo(o)
	c.exports = o
	return
}

// Closed return wheather the connection is closed by the script
func (c *WrappedConn)Closed()(bool){
	return c.closed.Load()
}

func (c *WrappedConn)Exports()(*goja.Object){
	return c.exports
}

func (c *WrappedConn)Emit(event *Event)(done <-chan bool){
	return c.emitter.EmitAsync(event)
}

func (c *WrappedConn)Recv()(w *WrappedPacketReader, err error){
	r, err := c.Conn.Recv()
	if err != nil {
		return
	}
	ch := make(chan *WrappedPacketReader, 1)
	c.loop.RunOnLoop(func(vm *goja.Runtime){
		ch <- WrapPacketReader(r, vm)
	})
	return <-ch, nil
}

// func (c *WrappedConn)Reset(conn *liter.Conn){
// 	c.Conn = conn
// }
