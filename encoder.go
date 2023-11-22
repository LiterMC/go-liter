
// This package provide some useful data encode/decode
package liter

import (
	"errors"
	"io"
	"math"
)

var VarIntTooBig  = errors.New("VarInt is too big")
var VarLongTooBig = errors.New("VarLong is too big")

func encodeUint16(buf []byte, v uint16)([]byte){
	buf[0] = (byte)(v >> 8 & 0xff)
	buf[1] = (byte)(v & 0xff)
	return buf
}

func encodeUint32(buf []byte, v uint32)([]byte){
	buf[0] = (byte)(v >> 24 & 0xff)
	buf[1] = (byte)(v >> 16 & 0xff)
	buf[2] = (byte)(v >> 8 & 0xff)
	buf[3] = (byte)(v & 0xff)
	return buf
}

func encodeUint64(buf []byte, v uint64)([]byte){
	buf[0] = (byte)(v >> 56 & 0xff)
	buf[1] = (byte)(v >> 48 & 0xff)
	buf[2] = (byte)(v >> 40 & 0xff)
	buf[3] = (byte)(v >> 32 & 0xff)
	buf[4] = (byte)(v >> 24 & 0xff)
	buf[5] = (byte)(v >> 16 & 0xff)
	buf[6] = (byte)(v >> 8 & 0xff)
	buf[7] = (byte)(v & 0xff)
	return buf
}

func encodeFloat32(buf []byte, v float32)([]byte){
	return encodeUint32(buf, math.Float32bits(v))
}

func encodeFloat64(buf []byte, v float64)([]byte){
	return encodeUint64(buf, math.Float64bits(v))
}


func decodeUint16(buf []byte)(uint16){
	return (uint16)(buf[0]) << 8 | (uint16)(buf[1])
}

func decodeUint32(buf []byte)(uint32){
	return (uint32)(buf[0]) << 24 | (uint32)(buf[1]) << 16 |
		(uint32)(buf[2]) << 8 | (uint32)(buf[3])
}

func decodeUint64(buf []byte)(uint64){
	return (uint64)(buf[0]) << 56 | (uint64)(buf[1]) << 48 |
		(uint64)(buf[2]) << 40 | (uint64)(buf[3]) << 32 |
		(uint64)(buf[4]) << 24 | (uint64)(buf[5]) << 16 |
		(uint64)(buf[6]) << 8 | (uint64)(buf[7])
}

func decodeFloat32(buf []byte)(float32){
	return math.Float32frombits(decodeUint32(buf))
}

func decodeFloat64(buf []byte)(float64){
	return math.Float64frombits(decodeUint64(buf))
}

func encodeVarInt[T ~int32](buf []byte, v T)(n int){
	w := (uint32)(v)
	for w &^ 0x7f != 0 {
		buf[n] = (byte)((w & 0x7f) | 0x80)
		n++
		w >>= 7
	}
	buf[n] = (byte)(w)
	n++
	return
}

func encodeVarLong[T ~int64](buf []byte, v T)(n int){
	w := (uint64)(v)
	for w &^ 0x7f != 0 {
		buf[n] = (byte)((w & 0x7f) | 0x80)
		n++
		w >>= 7
	}
	buf[n] = (byte)(w)
	n++
	return
}

func decodeVarInt(buf []byte)(n int, v VarInt){
	var (
		i int = 0
		b byte
		v0 uint32
	)
	for {
		if n >= len(buf) {
			return -1, 0
		}
		b = buf[n]
		n++
		v0 |= (uint32)(b & 0x7f) << i
		if b & 0x80 == 0 {
			return n, (VarInt)(v0)
		}
		if i += 7; i >= 32 {
			panic("VarInt is too big")
		}
	}
}

func decodeVarLong(buf []byte)(n int, v VarLong){
	var (
		i int = 0
		b byte
		v0 uint64
	)
	for {
		if n >= len(buf) {
			return -1, 0
		}
		b = buf[n]
		n++
		v0 |= (uint64)(b & 0x7f) << i
		if b & 0x80 == 0 {
			return n, (VarLong)(v0)
		}
		if i += 7; i >= 64 {
			panic(VarLongTooBig)
		}
	}
}

func readVarInt(r io.Reader)(n int, v int32, err error){
	var (
		i int = 0
		b [1]byte
		v0 uint32
	)
	for {
		if _, err = r.Read(b[:]); err != nil {
			return
		}
		n++
		v0 |= (uint32)(b[0] & 0x7f) << i
		if b[0] & 0x80 == 0 {
			v = (int32)(v0)
			return
		}
		if i += 7; i >= 32 {
			err = VarIntTooBig
			return
		}
	}
}

func readVarLong(r io.Reader)(n int, v int64, err error){
	var (
		i int = 0
		b [1]byte
		v0 uint64
	)
	for {
		if _, err = r.Read(b[:]); err != nil {
			return
		}
		n++
		v0 |= (uint64)(b[0] & 0x7f) << i
		if b[0] & 0x80 == 0 {
			v = (int64)(v0)
			return
		}
		if i += 7; i >= 64 {
			err = VarLongTooBig
			return
		}
	}
}
