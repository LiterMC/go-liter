package script

import (
	"io"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/google/uuid"
	"github.com/kmcsr/go-liter"
)

func WrapPacketBuilder(p *liter.PacketBuilder, conn *liter.Conn, vm *goja.Runtime, loop *eventloop.EventLoop) (o *goja.Object) {
	o = vm.NewObject()
	sent := false
	o.DefineAccessorProperty("protocol", vm.ToValue(func(goja.FunctionCall) goja.Value {
		return vm.ToValue(p.Protocol())
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	o.DefineAccessorProperty("id", vm.ToValue(func(goja.FunctionCall) goja.Value {
		return vm.ToValue((int)(p.Id()))
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	o.Set("send", func(call goja.FunctionCall) goja.Value {
		if sent {
			panic(vm.ToValue("This packet is already being sent"))
		}
		sent = true
		pm, r, e := vm.NewPromise()
		go func() {
			if err := conn.Send(p); err != nil {
				loop.RunOnLoop(func(*goja.Runtime) {
					e(err)
				})
			} else {
				loop.RunOnLoop(func(*goja.Runtime) {
					r(goja.Undefined())
				})
			}
		}()
		return vm.ToValue(pm)
	})
	o.Set("bytearray", func(call goja.FunctionCall) goja.Value {
		p.ByteArray(jsValue2Bytes(call.Arguments[0], vm))
		return o
	})
	o.Set("bool", func(call goja.FunctionCall) goja.Value {
		p.Bool(call.Arguments[0].ToBoolean())
		return o
	})
	o.Set("byte", func(call goja.FunctionCall) goja.Value {
		p.Byte((liter.Byte)(call.Arguments[0].ToInteger()))
		return o
	})
	o.Set("ubyte", func(call goja.FunctionCall) goja.Value {
		p.UByte((liter.UByte)(call.Arguments[0].ToInteger()))
		return o
	})
	o.Set("short", func(call goja.FunctionCall) goja.Value {
		p.Short((liter.Short)(call.Arguments[0].ToInteger()))
		return o
	})
	o.Set("ushort", func(call goja.FunctionCall) goja.Value {
		p.UShort((liter.UShort)(call.Arguments[0].ToInteger()))
		return o
	})
	o.Set("int", func(call goja.FunctionCall) goja.Value {
		p.Int((liter.Int)(call.Arguments[0].ToInteger()))
		return o
	})
	o.Set("long", func(call goja.FunctionCall) goja.Value {
		p.Long((liter.Long)(call.Arguments[0].ToInteger()))
		return o
	})
	o.Set("varInt", func(call goja.FunctionCall) goja.Value {
		p.VarInt((liter.VarInt)(call.Arguments[0].ToInteger()))
		return o
	})
	o.Set("varLong", func(call goja.FunctionCall) goja.Value {
		p.VarLong((liter.VarLong)(call.Arguments[0].ToInteger()))
		return o
	})
	o.Set("float", func(call goja.FunctionCall) goja.Value {
		p.Float((liter.Float)(call.Arguments[0].ToFloat()))
		return o
	})
	o.Set("double", func(call goja.FunctionCall) goja.Value {
		p.Double((liter.Double)(call.Arguments[0].ToFloat()))
		return o
	})
	o.Set("string", func(call goja.FunctionCall) goja.Value {
		p.String((liter.String)(call.Arguments[0].String()))
		return o
	})
	o.Set("uuid", func(call goja.FunctionCall) goja.Value {
		id, err := uuid.Parse(call.Arguments[0].String())
		if err != nil {
			panic(vm.NewGoError(err))
		}
		p.UUID(id)
		return o
	})
	o.Set("json", func(call goja.FunctionCall) goja.Value {
		p.JSON(call.Arguments[0])
		return o
	})
	return
}

type WrappedPacketReader struct {
	*liter.PacketReader
	exports *goja.Object
}

var _ Exportsable = (*WrappedPacketReader)(nil)

func WrapPacketReader(p *liter.PacketReader, vm *goja.Runtime) *WrappedPacketReader {
	EOF := vm.NewGoError(io.EOF)
	o := vm.NewObject()
	o.DefineAccessorProperty("protocol", vm.ToValue(func(goja.FunctionCall) goja.Value {
		return vm.ToValue(p.Protocol())
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	o.DefineAccessorProperty("id", vm.ToValue(func(goja.FunctionCall) goja.Value {
		return vm.ToValue((int)(p.Id()))
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	o.DefineAccessorProperty("offset", vm.ToValue(func(goja.FunctionCall) goja.Value {
		return vm.ToValue(p.Offset())
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	o.DefineAccessorProperty("remain", vm.ToValue(func(goja.FunctionCall) goja.Value {
		return vm.ToValue(p.Remain())
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	o.DefineAccessorProperty("data", vm.ToValue(func(goja.FunctionCall) goja.Value {
		return vm.ToValue(vm.NewArrayBuffer(p.Bytes()))
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	o.Set("bytearray", func(call goja.FunctionCall) goja.Value {
		buf := make([]byte, call.Arguments[0].ToInteger())
		if !p.ByteArray(buf) {
			panic(EOF)
		}
		jsUint8Array, _ := goja.AssertConstructor(vm.Get("Uint8Array"))
		res, err := jsUint8Array(vm.NewObject(), vm.ToValue(vm.NewArrayBuffer(buf)))
		if err != nil {
			panic(err)
		}
		return res
	})
	o.Set("bool", func(call goja.FunctionCall) goja.Value {
		v, ok := p.Bool()
		if !ok {
			panic(EOF)
		}
		return vm.ToValue(v)
	})
	o.Set("byte", func(call goja.FunctionCall) goja.Value {
		v, ok := p.Byte()
		if !ok {
			panic(EOF)
		}
		return vm.ToValue(v)
	})
	o.Set("ubyte", func(call goja.FunctionCall) goja.Value {
		v, ok := p.UByte()
		if !ok {
			panic(EOF)
		}
		return vm.ToValue(v)
	})
	o.Set("short", func(call goja.FunctionCall) goja.Value {
		v, ok := p.Short()
		if !ok {
			panic(EOF)
		}
		return vm.ToValue(v)
	})
	o.Set("ushort", func(call goja.FunctionCall) goja.Value {
		v, ok := p.UShort()
		if !ok {
			panic(EOF)
		}
		return vm.ToValue(v)
	})
	o.Set("int", func(call goja.FunctionCall) goja.Value {
		v, ok := p.Int()
		if !ok {
			panic(EOF)
		}
		return vm.ToValue(v)
	})
	o.Set("long", func(call goja.FunctionCall) goja.Value {
		v, ok := p.Long()
		if !ok {
			panic(EOF)
		}
		return vm.ToValue(v)
	})
	o.Set("varInt", func(call goja.FunctionCall) goja.Value {
		v, ok := p.VarInt()
		if !ok {
			panic(EOF)
		}
		return vm.ToValue(v)
	})
	o.Set("varLong", func(call goja.FunctionCall) goja.Value {
		v, ok := p.VarLong()
		if !ok {
			panic(EOF)
		}
		return vm.ToValue(v)
	})
	o.Set("float", func(call goja.FunctionCall) goja.Value {
		v, ok := p.Float()
		if !ok {
			panic(EOF)
		}
		return vm.ToValue(v)
	})
	o.Set("double", func(call goja.FunctionCall) goja.Value {
		v, ok := p.Double()
		if !ok {
			panic(EOF)
		}
		return vm.ToValue(v)
	})
	o.Set("string", func(call goja.FunctionCall) goja.Value {
		v, ok := p.String()
		if !ok {
			panic(EOF)
		}
		return vm.ToValue(v)
	})
	o.Set("uuid", func(call goja.FunctionCall) goja.Value {
		id, ok := p.UUID()
		if !ok {
			panic(EOF)
		}
		return vm.ToValue(id.String())
	})
	o.Set("json", func(call goja.FunctionCall) goja.Value {
		var v any
		if err := p.JSON(&v); err != nil {
			panic(vm.NewGoError(err))
		}
		return vm.ToValue(v)
	})
	return &WrappedPacketReader{
		PacketReader: p,
		exports:      o,
	}
}

func (p *WrappedPacketReader) Exports() *goja.Object {
	return p.exports
}

func jsValue2Bytes(v goja.Value, vm *goja.Runtime) (b []byte) {
	switch w := v.Export().(type) {
	case goja.ArrayBuffer:
		b = w.Bytes()
	case []byte:
		b = w
	default:
		v := v.ToObject(vm)
		jsArray := vm.Get("Array").ToObject(vm)
		jsUint8Array := vm.Get("Uint8Array").ToObject(vm)
		switch {
		case vm.InstanceOf(v, jsArray):
			newer, ok := goja.AssertConstructor(jsUint8Array)
			if !ok {
				panic("jsUint8Array is not a constructor")
			}
			var err error
			if v, err = newer(nil, v); err != nil {
				panic(err)
			}
			fallthrough
		case vm.InstanceOf(v, jsUint8Array):
			b = v.Get("buffer").Export().(goja.ArrayBuffer).Bytes()
		default:
			panic(vm.NewTypeError("Unknown type of arg#0"))
		}
	}
	return
}
