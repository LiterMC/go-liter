package nbt

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"

	liter "github.com/kmcsr/go-liter"
)

type (
	PacketBuilder = liter.PacketBuilder
	PacketReader  = liter.PacketReader
	Byte          = liter.Byte
	UByte         = liter.UByte
	Short         = liter.Short
	UShort        = liter.UShort
	Int           = liter.Int
	Long          = liter.Long
	Float         = liter.Float
	Double        = liter.Double
	String        = liter.String
	ByteArray     = liter.ByteArray
)

type NBTIdNotExistsErr struct {
	Id Byte
}

func (e *NBTIdNotExistsErr) Error() string {
	return fmt.Sprintf("NBT id 0x%02x is not exists", e.Id)
}

type NBT interface {
	Type() Byte
	Name() string
	SetName(name string)

	fmt.Stringer
	liter.Encodable
	liter.Decodable
}

const (
	NbtUnknown   = -1
	NbtEnd       = 0x00
	NbtByte      = 0x01
	NbtShort     = 0x02
	NbtInt       = 0x03
	NbtLong      = 0x04
	NbtFloat     = 0x05
	NbtDouble    = 0x06
	NbtByteArray = 0x07
	NbtString    = 0x08
	NbtList      = 0x09
	NbtCompound  = 0x0a
	NbtIntArray  = 0x0b
	NbtLongArray = 0x0c
)

var NBTNewer = map[Byte]func() NBT{
	NbtEnd:       func() NBT { return _NBTEndIns },
	NbtByte:      func() NBT { return new(NBTByte) },
	NbtShort:     func() NBT { return new(NBTShort) },
	NbtInt:       func() NBT { return new(NBTInt) },
	NbtLong:      func() NBT { return new(NBTLong) },
	NbtFloat:     func() NBT { return new(NBTFloat) },
	NbtDouble:    func() NBT { return new(NBTDouble) },
	NbtByteArray: func() NBT { return new(NBTByteArray) },
	NbtString:    func() NBT { return new(NBTString) },
	NbtList:      func() NBT { return new(NBTList) },
	NbtCompound:  func() NBT { return new(NBTCompound) },
	NbtIntArray:  func() NBT { return new(NBTIntArray) },
	NbtLongArray: func() NBT { return new(NBTLongArray) },
}

type Decompressor = func(in []byte) (out []byte, ok bool)

var dataDecompressors = []Decompressor{
	func(in []byte) (out []byte, ok bool) { // gzip
		r, err := gzip.NewReader(bytes.NewReader(in))
		if err != nil {
			return nil, false
		}
		defer r.Close()
		out, err = io.ReadAll(r)
		if err != nil {
			return nil, false
		}
		return out, true
	},
	func(in []byte) (out []byte, ok bool) { // zlib
		r, err := zlib.NewReader(bytes.NewReader(in))
		if err != nil {
			return nil, false
		}
		defer r.Close()
		out, err = io.ReadAll(r)
		if err != nil {
			return nil, false
		}
		return out, true
	},
}

func WriteNBT(b *PacketBuilder, n NBT) {
	b.Byte(n.Type())
	writeNBTString(b, n.Name())
	n.Encode(b)
}

func ReadNBT(r *PacketReader) (n NBT, err error) {
	var id Byte
	var ok bool
	if id, ok = r.Byte(); !ok {
		return nil, io.EOF
	}
	if id == NbtEnd {
		return _NBTEndIns, nil
	}
	var newer func() NBT
	if newer, ok = NBTNewer[id]; !ok {
		return nil, &NBTIdNotExistsErr{id}
	}
	n = newer()
	var name string
	if name, ok = readNBTString(r); !ok {
		return nil, io.EOF
	}
	n.SetName(name)
	n.DecodeFrom(r)
	return
}

func ParseFromBytes(buf []byte) (nbt NBT, err error) {
	var data []byte
	var ok bool
	for _, dc := range dataDecompressors {
		if data, ok = dc(buf); ok {
			break
		}
	}
	if !ok {
		data = buf
	}
	r := liter.ReadPacketFromBytes(liter.V_UNSET, data)
	return ReadNBT(r)
}
