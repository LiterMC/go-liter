
package script

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/dop251/goja"
	"github.com/dop251/goja/ast"
	"github.com/dop251/goja/file"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/require"
	"github.com/kmcsr/go-logger"
	"github.com/kmcsr/go-logger/logrus"
	"github.com/kmcsr/go-liter/script/console"
)

var (
	fileNameRe = regexp.MustCompile(`^([a-z_][0-9a-z_]{0,31})(?:@(\d+(?:\.\d+)*)(?:-.+)?)?\..+$`)
)

var (
	ErrFileNameInvalid = errors.New("The script's filename is invalid")
	ErrScriptLoaded    = errors.New("The script is already loaded")
	ErrScriptInvalid   = errors.New("Invalid script")
)

type Manager struct {
	logger logger.Logger

	loop *eventloop.EventLoop
	scriptMux sync.Mutex
	scripts   map[string]*Script
}

func NewManager()(m *Manager){
	m = &Manager{
		logger: logrus.Logger,
		loop: eventloop.NewEventLoop(
			eventloop.WithRegistry(require.NewRegistryWithLoader(disableRequire)),
			eventloop.EnableConsole(false)),
		scripts: make(map[string]*Script),
	}
	m.loop.Start()
	m.loop.RunOnLoop(func(vm *goja.Runtime){
		vm.GlobalObject().Delete("require") // disable require
	})
	return
}

func (m *Manager)SetLogger(loger logger.Logger){
	m.scriptMux.Lock()
	defer m.scriptMux.Unlock()
	m.logger = loger
	for _, s := range m.scripts {
		s.console.SetLogger(setPrefixLogger(loger, s.id))
	}
}

var errRequireDisabled = errors.New("require is disabled")

func disableRequire(path string)([]byte, error){
	return nil, errRequireDisabled
}

func (m *Manager)List()(scripts []*Script){
	m.scriptMux.Lock()
	defer m.scriptMux.Unlock()

	scripts = make([]*Script, 0, len(m.scripts))
	for _, s := range m.scripts {
		scripts = append(scripts, s)
	}
	return
}

func (m *Manager)Load(path string)(script *Script, err error){
	return m.LoadWithContext(context.Background(), path)
}

func (m *Manager)LoadWithContext(ctx context.Context, path string)(script *Script, err error){
	name := filepath.Base(path)
	matches := fileNameRe.FindStringSubmatch(name)
	if matches == nil {
		return nil, ErrFileNameInvalid
	}
	id := matches[1]
	version := matches[2]

	m.scriptMux.Lock()
	if _, ok := m.scripts[id]; ok {
		m.scriptMux.Unlock()
		return nil, ErrScriptLoaded
	}
	m.scripts[id] = nil // reserve the slot
	m.scriptMux.Unlock()

	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	parsed, err := goja.Parse(path, (string)(data))
	if err != nil {
		return
	}
	// wrap code as `function($, console){ ... }`
	parsed.Body = []ast.Statement{&ast.ExpressionStatement{
		Expression: &ast.FunctionLiteral{
			Function: -1,
			Name: nil,
			ParameterList: &ast.ParameterList{
				Opening: -1,
				List: []*ast.Binding{
					{ Target: &ast.Identifier{ Name:"$", Idx: -1 } },
					{ Target: &ast.Identifier{ Name:"console", Idx: -1 } },
				},
				Closing: -1,
			},
			Body: &ast.BlockStatement{
				LeftBrace: 0,
				List: parsed.Body,
				RightBrace: (file.Idx)(len(data) - 1),
			},
			DeclarationList: parsed.DeclarationList,
		},
	}}
	prog, err := goja.CompileAST(parsed, true)
	if err != nil {
		return
	}

	script = newScript(id, version, path, prog)

	errCh := make(chan error, 1)
	m.loop.RunOnLoop(func(vm *goja.Runtime){
		defer close(errCh)
		var err error
		var res goja.Value
		if res, err = vm.RunProgram(prog); err != nil {
			errCh <- err
			return
		}
		cb, ok := goja.AssertFunction(res)
		if !ok {
			errCh <- ErrScriptInvalid
			return
		}
		script.init(vm)
		script.console = console.NewConsole(vm, setPrefixLogger(m.logger, id))
		if _, err = cb(nil, script.doll, script.console.Exports()); err != nil {
			errCh <- err
			return
		}
	})
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err = <-errCh:
		if err != nil {
			return
		}
	}
	m.scriptMux.Lock()
	defer m.scriptMux.Unlock()
	m.scripts[id] = script
	return
}

func setPrefixLogger(base logger.Logger, prefix string)(logger.Logger){
	// if l, ok := base.(*logrus.LogrusWrapper); ok {
	// 	return l.WithFields(logrus.Fields{"prefix": prefix})
	// }
	return logger.NewPrefixLogger(base, "[" + prefix + "]")
}

func (m *Manager)Unload(id string)(script *Script){
	m.scriptMux.Lock()
	defer m.scriptMux.Unlock()
	if script = m.scripts[id]; script == nil {
		return
	}
	delete(m.scripts, id)
	return
}

// LoadFromDir will load all scripts under the path which have the ext `.js`
// If the target path is not exists, LoadFromDir will do nothing and return no error
func (m *Manager)LoadFromDir(path string)(scripts []*Script, err error){
	entries, err := os.ReadDir(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = nil
		}
		return
	}
	errs := make([]error, 0, len(entries) / 2)
	for _, e := range entries {
		name := e.Name()
		if !e.IsDir() {
			if strings.HasSuffix(name, ".js") {
				s, err := m.Load(filepath.Join(path, name))
				if err != nil {
					errs = append(errs, err)
					continue
				}
				scripts = append(scripts, s)
			}
		}
	}
	if len(errs) > 0 {
		err = errors.Join(errs...)
	}
	return
}
