
package main

import (
	"bytes"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/google/uuid"
)

func split(s string, chs string)(a, b string){
	i := strings.IndexAny(s, chs)
	if i == -1 {
		return s, ""
	}
	return s[:i], s[i + 1:]
}

func ismatch(name string, pattern string)(bool){
	if len(pattern) > 0 && pattern[0] == '*' {
		return len(pattern) == 1 || strings.HasSuffix(name, pattern[1:])
	}
	return name == pattern
}

func waitUntilNot(cond *sync.Cond, not func()(bool))(<-chan struct{}){
	ch := make(chan struct{}, 0)
	go func(){
		defer close(ch)
		cond.L.Lock()
		defer cond.L.Unlock()
		for not() {
			cond.Wait()
		}
	}()
	return ch
}

func matchIP(pattern string, ip net.IP) bool {
	pattern, a := split(pattern, "/")
	var n int
	if len(a) > 0 {
		n, _ = strconv.Atoi(a)
	}
	pip := net.ParseIP(pattern)
	if n > 0 {
		if len(pip) != len(ip) {
			pip, ip = pip.To16(), ip.To16()
		}
		l := len(ip)
		e, m := l - n / 8, (byte)(0xff) << (n % 8)
		if m != 0 {
			e--
		}
		if bytes.Equal(pip[:e], ip[:e]) && (pip[e] & m == ip[e] & m) {
			return true
		}
		return false
	}
	return pip.Equal(ip)
}

func uuidLess(a, b uuid.UUID)(bool){
	for i, n := range a {
		if n < b[i] {
			return true
		}
	}
	return false
}
