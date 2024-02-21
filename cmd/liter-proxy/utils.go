package main

import (
	"sync"
)

type Map = map[string]any

type Set[T comparable] struct {
	mux  sync.RWMutex
	cond *sync.Cond
	m    map[T]struct{}
}

func NewSet[T comparable]() (s *Set[T]) {
	s = &Set[T]{
		m: make(map[T]struct{}),
	}
	s.cond = sync.NewCond(&s.mux)
	return
}

func (s *Set[T]) Add(v T) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m[v] = struct{}{}
	s.cond.Broadcast()
}

func (s *Set[T]) Has(v T) (ok bool) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	_, ok = s.m[v]
	return
}

func (s *Set[T]) Del(v T) (ok bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if _, ok = s.m[v]; ok {
		delete(s.m, v)
	}
	s.cond.Broadcast()
	return
}

func (s *Set[T]) Len() (n int) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return len(s.m)
}

func (s *Set[T]) AsSlice() (sli []T) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	sli = make([]T, 0, len(s.m))
	for v, _ := range s.m {
		sli = append(sli, v)
	}
	return
}

func (s *Set[T]) AfterEmpty() (ch <-chan struct{}) {
	c := make(chan struct{}, 0)
	go func() { // maybe goroutine leak here if we alloc many sets
		defer close(c)
		for s.Len() != 0 {
			s.cond.Wait()
		}
	}()
	return c
}
