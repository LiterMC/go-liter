
package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"net"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

func remove[T any](slice []T, index int)([]T){
	copy(slice[index:], slice[index + 1:])
	return slice[:len(slice) - 1]
}

func insert[T any](slice []T, index int, args ...T)(res []T){
	n := len(args)
	if n == 0 {
		return slice
	}
	if index >= len(slice) {
		return append(slice, args...)
	}
	tmp := make([]T, len(slice))
	copy(tmp, slice)
	resLen := len(slice) + n
	if ca := cap(slice); ca >= resLen {
		slice = slice[:resLen]
	}else{
		if ca == 0 {
			ca = resLen
		}else{
			for ca < resLen {
				if ca < 1024 {
					ca *= 2
				}else{
					ca += 1024
				}
			}
		}
		slice = make([]T, resLen, ca)
	}
	copy(slice, tmp[:index])
	copy(slice[index:], args)
	copy(slice[index + n:], tmp[index:])
	return slice
}

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

func genRandB64(n int)(s string, err error){
	buf := make([]byte, n)
	if _, err = rand.Read(buf); err != nil {
		return
	}
	s = base64.RawURLEncoding.EncodeToString(buf)
	return
}

func asSha256(s string)(string){
	buf := sha256.Sum256(([]byte)(s))
	return base64.RawURLEncoding.EncodeToString(buf[:])
}

func toSeconds(t time.Time)(float64){
	return (float64)(t.Unix()) + (float64)(t.UnixNano() % 1e9) / 1e9
}

var stackPool = sync.Pool{
	New: func()(any){
		buf := make([]byte, 1024)
		return &buf
	},
}

func getStacktrace()(s string){
	buf := *(stackPool.Get().(*[]byte))
	defer func(){
		stackPool.Put(&buf)
	}()
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return (string)(buf[:n])
		}
		stackPool.Put(&buf)
		buf = make([]byte, 2 * len(buf))
	}
}
