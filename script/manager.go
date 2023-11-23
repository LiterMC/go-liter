package script

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/require"
	"github.com/kmcsr/go-liter"
	"github.com/kmcsr/go-logger"
	"github.com/kmcsr/go-logger/logrus"
)

type PluginIdLoadedError struct {
	Id string
}

func (e *PluginIdLoadedError) Error() string {
	return fmt.Sprintf("Plugin %s is already loaded", e.Id)
}

type PluginLoadError struct {
	Path   string
	Origin error
}

func (e *PluginLoadError) Unwrap() error {
	return e.Origin
}

func (e *PluginLoadError) Error() string {
	return fmt.Sprintf("Error when loading %q: %v", e.Path, e.Origin)
}

type Manager struct {
	logger logger.Logger

	loop      *eventloop.EventLoop
	scriptMux sync.RWMutex
	scripts   map[string]*Script
}

func NewManager() (m *Manager) {
	m = &Manager{
		logger: logrus.Logger,
		loop: eventloop.NewEventLoop(
			eventloop.WithRegistry(require.NewRegistryWithLoader(disableRequire)),
			eventloop.EnableConsole(false)),
		scripts: make(map[string]*Script),
	}
	m.loop.Start()
	m.loop.RunOnLoop(func(vm *goja.Runtime) {
		vm.GlobalObject().Delete("require") // disable the default require provided by goja_nodejs
	})
	return
}

func (m *Manager) SetLogger(loger logger.Logger) {
	m.scriptMux.Lock()
	defer m.scriptMux.Unlock()
	m.logger = loger
	for _, s := range m.scripts {
		s.console.SetLogger(setPrefixLogger(loger, s.Id))
	}
}

var errRequireDisabled = errors.New("the default require is disabled")

func disableRequire(path string) ([]byte, error) {
	return nil, errRequireDisabled
}

// List will return all active plugins
func (m *Manager) List() (scripts []*Script) {
	m.scriptMux.Lock()
	defer m.scriptMux.Unlock()

	scripts = make([]*Script, 0, len(m.scripts))
	for _, s := range m.scripts {
		if s != nil {
			scripts = append(scripts, s)
		}
	}
	return
}

// Load will call LoadWithContext with context.Background()
func (m *Manager) Load(path string) (script *Script, err error) {
	return m.LoadWithContext(context.Background(), path)
}

// LoadWithContext will load a plugin packet use the given filepath.
// The first capture group will be the script's ID. The second capture group is the script's version
// If the script's ID is already loaded, then an error will be returned
func (m *Manager) LoadWithContext(ctx context.Context, path string) (script *Script, err error) {
	var packet *zip.ReadCloser
	if packet, err = zip.OpenReader(path); err != nil {
		return
	}
	meta, err := loadScriptMeta(packet)
	if err != nil {
		return
	}
	m.scriptMux.Lock()
	if _, ok := m.scripts[meta.Id]; ok {
		m.scriptMux.Unlock()
		return nil, &PluginIdLoadedError{Id: meta.Id}
	}
	m.scripts[meta.Id] = nil // reserve the slot
	m.scriptMux.Unlock()

	errCh := make(chan error, 1)
	m.loop.RunOnLoop(func(vm *goja.Runtime) {
		var err error
		defer func() {
			errCh <- err
		}()
		if script, err = loadScript(packet, meta, m.extModuleLoader, m.logger, vm, m.loop); err != nil {
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
	m.scripts[script.Id] = script
	return
}

func setPrefixLogger(base logger.Logger, prefix string) logger.Logger {
	// if l, ok := base.(*logrus.LogrusWrapper); ok {
	// 	return l.WithFields(logrus.Fields{"prefix": prefix})
	// }
	return logger.NewPrefixLogger(base, "["+prefix+"]")
}

// Unload will unload plugin by ID and return the unloaded plugin instance
// During unloading, the unload event will be emitted
func (m *Manager) Unload(id string) (script *Script) {
	m.scriptMux.Lock()
	defer m.scriptMux.Unlock()

	if script = m.scripts[id]; script == nil {
		return
	}
	<-script.Emit(&Event{
		Name: "unload",
	})
	script.loop = nil
	delete(m.scripts, id)
	return
}

// UnloadAll will unload all plugins and return them
// During unloading, the unload event will be emitted
func (m *Manager) UnloadAll() (scripts []*Script) {
	m.scriptMux.Lock()
	defer m.scriptMux.Unlock()

	if len(m.scripts) == 0 {
		return
	}

	scripts = make([]*Script, 0, len(m.scripts))
	dones := make([]<-chan bool, 0, len(m.scripts))
	for _, s := range m.scripts {
		if s != nil {
			scripts = append(scripts, s)
			dones = append(dones, s.Emit(&Event{
				Name: "unload",
			}))
		}
	}
	for _, ch := range dones {
		<-ch
	}
	return
}

// LoadFromDir will load all plugins under the path which have the ext `.zip`,
// and return all successfully loaded plugins with a possible error
// If the target path is not exists, LoadFromDir will do nothing and return without error
// If there are errors during load any plugin, the errors will be wrapped use `errors.Join`,
// and other plugins will continue to be load.
func (m *Manager) LoadFromDir(path string) (scripts []*Script, err error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = nil
		}
		return
	}
	errs := make([]error, 0, len(entries)/2)
	for _, e := range entries {
		name := e.Name()
		if !e.IsDir() {
			if strings.HasSuffix(name, ".zip") {
				p := filepath.Join(path, name)
				s, err := m.Load(p)
				if err != nil {
					errs = append(errs, &PluginLoadError{Path: p, Origin: err})
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

func (m *Manager) extModuleLoader(module string) (l *moduleLoader, err error) {
	s := m.scripts[module]
	if s != nil {
		return s.loader, nil
	}
	return nil, &ModuleNotFoundErr{
		Module: module,
	}
}

func (m *Manager) Emit(event *Event) (done <-chan bool) {
	if event.Cancelled() {
		panic("liter: script: Trying to emit a cancelled event")
	}

	exit := make(chan bool, 1)
	go func() {
		defer close(exit)
		m.scriptMux.RLock()
		defer m.scriptMux.RUnlock()

		for _, s := range m.scripts {
			if s != nil {
				if <-s.Emit(event) {
					exit <- true
					return
				}
			}
		}
	}()
	return exit
}

func (m *Manager) WrapConn(conn *liter.Conn) (wrapped *WrappedConn) {
	done := make(chan struct{}, 0)
	m.loop.RunOnLoop(func(vm *goja.Runtime) {
		defer close(done)
		wrapped = WrapConn(conn, vm, m.loop)
	})
	<-done
	return
}
