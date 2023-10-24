
package script

import (
	"errors"
	"io"
	"io/fs"
	stdpath "path"

	"github.com/dop251/goja"
	"github.com/dop251/goja/ast"
	"github.com/dop251/goja/file"
	"github.com/dop251/goja/parser"
	"github.com/dop251/goja/unistring"
)

var (
	ErrScriptInvalid  = errors.New("Invalid script")
	ErrCircularImport = errors.New("Circular import")
	ErrIsDir          = errors.New("Target file is a directory")
)

type addonVar struct {
	name string
	val  goja.Value
}

type moduleLoader struct {
	vm *goja.Runtime
	packet fs.FS
	loading map[string]struct{}
	loaded map[string]*goja.Object
	addonVars []goja.Value
	paramAst  []*ast.Binding
}

func newModuleLoader(packet fs.FS, vm *goja.Runtime, vars []addonVar)(r *moduleLoader){
	paramAst := make([]*ast.Binding, 3 + len(vars))
	paramAst[0] = &ast.Binding{ Target: &ast.Identifier{ Name:"require", Idx: -1 } }
	paramAst[1] = &ast.Binding{ Target: &ast.Identifier{ Name:"module", Idx: -1 } }
	paramAst[2] = &ast.Binding{ Target: &ast.Identifier{ Name:"exports", Idx: -1 } }
	addonVars := make([]goja.Value, len(vars))
	for i, v := range vars {
		paramAst[i + 3] = &ast.Binding{
			Target: &ast.Identifier{ Name: unistring.NewFromString(v.name), Idx: -1 },
		}
		addonVars[i] = v.val
	}
	return &moduleLoader{
		vm: vm,
		packet: packet,
		loading: make(map[string]struct{}, 3),
		loaded: make(map[string]*goja.Object, 3),
		addonVars: addonVars,
		paramAst: paramAst,
	}
}

func (r *moduleLoader)makeRequire(base string)(goja.Value){
	return r.vm.ToValue(func(call goja.FunctionCall)(goja.Value){
		path := call.Argument(0).String()
		res, err := r.load(path, base)
		if err != nil {
			panic(wrap2JsErr(r.vm, err))
		}
		return res
	})
}

func (r *moduleLoader)load(path string, base string)(exports *goja.Object, err error){
	fd, path, err := r.resolveAndOpen(path, base)
	if err != nil {
		return
	}
	defer fd.Close()
	exports, ok := r.loaded[path]
	if ok {
		return
	}
	if _, ok = r.loading[path]; ok {
		return nil, ErrCircularImport
	}
	r.loading[path] = struct{}{}
	defer delete(r.loading, path)

	data, err := io.ReadAll(fd)
	if err != nil {
		return
	}
	parsed, err := goja.Parse(path, (string)(data), parser.WithSourceMapLoader(func(name string)(data []byte, err error){
		fd, err := r.packet.Open(name)
		if err != nil {
			return
		}
		defer fd.Close()
		return io.ReadAll(fd)
	}))
	if err != nil {
		return
	}

	parsed.Body = []ast.Statement{&ast.ExpressionStatement{
		Expression: &ast.FunctionLiteral{
			Function: -1,
			Name: nil,
			ParameterList: &ast.ParameterList{
				Opening: -1,
				List: r.paramAst,
				Closing: -1,
			},
			Body: &ast.BlockStatement{
				LeftBrace: 0,
				List: parsed.Body,
				RightBrace: (file.Idx)(len(parsed.File.Source())),
			},
			DeclarationList: parsed.DeclarationList,
		},
	}}
	prog, err := goja.CompileAST(parsed, true)
	if err != nil {
		return
	}
	var cb0 goja.Value
	if cb0, err = r.vm.RunProgram(prog); err != nil {
		return
	}
	cb, ok := goja.AssertFunction(cb0)
	if !ok {
		return nil, ErrScriptInvalid
	}

	module := r.vm.NewObject()
	module.Set("exports", r.vm.NewObject())

	args := make([]goja.Value, 3 + len(r.addonVars))
	args[0] = r.makeRequire(stdpath.Dir(path))
	args[1] = module
	args[2] = module.Get("exports")
	copy(args[3:], r.addonVars)
	if _, err = cb(nil, args...); err != nil {
		return
	}
	exports = module.Get("exports").ToObject(r.vm)
	r.loaded[path] = exports
	return
}

func openFileStat(fs fs.FS, path string)(fd fs.File, info fs.FileInfo, err error){
	if fd, err = fs.Open(path); err != nil {
		return
	}
	if info, err = fd.Stat(); err != nil {
		fd.Close()
		return
	}
	return
}

func (r *moduleLoader)resolveAndOpen(path string, base string)(fd fs.File, _ string, err error){
	path = stdpath.Join(base, path)
	var info fs.FileInfo
	var er error
	if fd, info, err = openFileStat(r.packet, path); err == nil {
		if !info.IsDir() {
			return fd, path, nil
		}
		p := path + "index.js"
		if fd, info, err = openFileStat(r.packet, p); err == nil {
			if info.IsDir() {
				err = ErrIsDir
				return
			}
			return fd, p, nil
		}
		p = path + "index.json"
		if fd, info, err = openFileStat(r.packet, p); err == nil {
			if info.IsDir() {
				err = ErrIsDir
				return
			}
			return fd, p, nil
		}
		return
	}
	if !errors.Is(err, fs.ErrNotExist) {
		return
	}
	if stdpath.Ext(path) != ".js" {
		p := path + ".js"
		if fd, info, er = openFileStat(r.packet, p); er == nil {
			if info.IsDir() {
				return
			}
			return fd, p, nil
		}
	}
	if stdpath.Ext(path) != ".json" {
		p := path + ".js"
		if fd, info, er = openFileStat(r.packet, p); er == nil {
			if info.IsDir() {
				return
			}
			return fd, p, nil
		}
	}
	return
}

func wrap2JsErr(vm *goja.Runtime, err error)(any){
	if _, ok := err.(*goja.Exception); ok {
		return err
	}
	return vm.NewGoError(err)
}
