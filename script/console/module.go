package console

import (
	"strconv"
	"strings"

	"github.com/dop251/goja"
	"github.com/kmcsr/go-logger"
)

type Console struct {
	runtime *goja.Runtime
	logger  logger.Logger
	exports *goja.Object
}

func NewConsole(runtime *goja.Runtime, loger logger.Logger) (c *Console) {
	c = &Console{
		runtime: runtime,
		logger:  loger,
	}
	o := runtime.NewObject()
	o.Set("trace", c.log(func(s string) { c.logger.Trace(s) }))
	o.Set("debug", c.log(func(s string) { c.logger.Debug(s) }))
	o.Set("log", c.log(func(s string) { c.logger.Info(s) }))
	o.Set("info", c.log(func(s string) { c.logger.Info(s) }))
	o.Set("warn", c.log(func(s string) { c.logger.Warn(s) }))
	o.Set("error", c.log(func(s string) { c.logger.Error(s) }))
	c.exports = o
	return
}

func (c *Console) Logger() logger.Logger {
	return c.logger
}

func (c *Console) SetLogger(loger logger.Logger) {
	c.logger = loger
}

func (c *Console) Exports() *goja.Object {
	return c.exports
}

func (c *Console) format(args ...goja.Value) string {
	if len(args) == 0 {
		return ""
	}
	var b strings.Builder
	argi := 1
	format := args[0].String()
	if _, ok := args[0].(goja.String); ok {
		var buf [16]byte
		b.Grow(len(format))
		formatting := false
		for _, c := range format {
			if formatting {
				formatting = false
				if c == '%' {
					b.WriteByte('%')
					continue
				}
				if argi >= len(args) {
					b.WriteByte('%')
					b.WriteRune(c)
					continue
				}
				switch c {
				case 't':
					b.Write(strconv.AppendBool(buf[:0], args[argi].ToBoolean()))
				case 'd', 'i':
					b.Write(strconv.AppendInt(buf[:0], args[argi].ToInteger(), 10))
				case 'x':
					b.Write(strconv.AppendInt(buf[:0], args[argi].ToInteger(), 16))
				case 'f':
					b.Write(strconv.AppendFloat(buf[:0], args[argi].ToFloat(), 'f', -1, 64))
				case 'o', 'q':
					arg := args[argi]
					if _, ok := arg.(goja.String); ok {
						b.Write(strconv.AppendQuote(buf[:0], arg.String()))
					} else {
						encode(&b, arg)
					}
				case 's':
					b.WriteString(args[argi].String())
				default:
					b.WriteByte('%')
					b.WriteRune(c)
					continue
				}
				argi++
			} else if c == '%' {
				formatting = true
			} else {
				b.WriteRune(c)
			}
		}
	} else {
		b.WriteString(format)
	}
	for _, v := range args[argi:] {
		b.WriteByte(' ')
		encode(&b, v)
	}
	return b.String()
}

func (c *Console) log(printer func(string)) func(goja.FunctionCall) goja.Value {
	return func(call goja.FunctionCall) (_ goja.Value) {
		printer(c.format(call.Arguments...))
		return
	}
}
