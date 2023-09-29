
package script

import (
	"sync"

	"github.com/dop251/goja"
)

type Storage interface {
	Len()(int)
	Keys()(keys []string)
	GetItem(key string)(value goja.Value)
	SetItem(key string, value goja.Value)
	RemoveItem(key string)
	Clear()
}

func ProxyStorage(s Storage, vm *goja.Runtime)(goja.Proxy){
	keys := vm.ToValue(s.Keys)
	getitem := vm.ToValue(s.GetItem)
	setitem := vm.ToValue(s.SetItem)
	removeitem := vm.ToValue(s.RemoveItem)
	clear := vm.ToValue(s.Clear)
	return vm.NewProxy(vm.NewObject(), &goja.ProxyTrapConfig{
		IsExtensible: func(_ *goja.Object)(success bool){
			return true
		},
		Has: func(_ *goja.Object, property string)(available bool){
			switch property {
			case "length", "keys", "getItem", "setItem", "removeItem", "clear":
				return true
			}
			return s.GetItem(property) != nil
		},
		Get: func(_ *goja.Object, property string, receiver goja.Value)(value goja.Value){
			switch property {
			case "length":
				return vm.ToValue(s.Len())
			case "keys":
				return keys
			case "getItem":
				return getitem
			case "setItem":
				return setitem
			case "removeItem":
				return removeitem
			case "clear":
				return clear
			}
			return s.GetItem(property)
		},
		Set: func(_ *goja.Object, property string, value goja.Value, receiver goja.Value)(success bool){
			s.SetItem(property, value)
			return true
		},
		DeleteProperty: func(_ *goja.Object, property string)(success bool){
			s.RemoveItem(property)
			return true
		},
		OwnKeys: func(_ *goja.Object)(object *goja.Object){
			return vm.ToValue(s.Keys()).ToObject(vm)
		},
	})
}

// MemoryStorage is designed for single thread use, so it's not thread-safe.
// It should only be called inside the js loop with same Runtime instance.
type MemoryStorage struct {
	data map[string]goja.Value
	exports *goja.Object
}

var _ Storage = (*MemoryStorage)(nil)

func NewMemoryStorage(vm *goja.Runtime)(s *MemoryStorage){
	s = &MemoryStorage{
		data: make(map[string]goja.Value),
	}
	s.exports = vm.ToValue(ProxyStorage(s, vm)).ToObject(vm)
	return
}

func (s *MemoryStorage)Exports()(*goja.Object){
	return s.exports
}

func (s *MemoryStorage)Len()(int){
	return len(s.data)
}

func (s *MemoryStorage)Keys()(keys []string){
	keys = make([]string, 0, len(s.data))
	for k, _ := range s.data {
		keys = append(keys, k)
	}
	return
}

func (s *MemoryStorage)GetItem(key string)(value goja.Value){
	return s.data[key]
}

func (s *MemoryStorage)SetItem(key string, value goja.Value){
	s.data[key] = value
}

func (s *MemoryStorage)RemoveItem(key string){
	delete(s.data, key)
}

func (s *MemoryStorage)Clear(){
	s.data = make(map[string]goja.Value)
}

// TODO
// JsonStorage is a local storage, it's thready-safe
type JsonStorage struct {
	path string
	mux  sync.RWMutex
	data map[string]goja.Value
}

// var _ Storage = (*JsonStorage)(nil)

// func NewJsonStorage(path string)(s *JsonStorage){
// 	s = &JsonStorage{
// 		path: path,
// 		data: make(map[string]goja.Value),
// 	}
// 	return
// }
