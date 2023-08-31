
package liter

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type (
	Bool   = bool
	Byte   = int8
	Short  = int16
	Int    = int32
	Long   = int64
	Float  = float32
	Double = float64

	VarInt  int32
	VarLong int64

	ByteArray = []byte
	String    = string
	Object    = map[string]any
	Position  struct {
		X int
		Y int
		Z int
	}
	UUID      = uuid.UUID

	Encodable interface {
		Encode(p *PacketBuilder)(err error)
	}
	Decodable interface {
		DecodeFrom(r *PacketReader)(err error)
	}
	Packet interface {
		Encodable
		Decodable
	}
)

func PositionFromLong(n Long)(p Position){
	return Position{
		X: (int)(n >> 38),
		Y: (int)(n << 52 >> 52),
		Z: (int)(n << 26 >> 38),
	}
}

var _ Encodable = Position{}

func (p Position)AsLong()(n Long){
	return (((Long)(p.X) & 0x3ffffff) << 38) | (((Long)(p.Z) & 0x3ffffff) << 12) | ((Long)(p.Y) & 0xfff)
}

func (p Position)Encode(b *PacketBuilder)(err error){
	b.Long(p.AsLong())
	return nil
}

type PacketBuilder struct{
	id int32
	buf []byte
	len int
}

func NewPacket(id int32)(p *PacketBuilder){
	return &PacketBuilder{
		id: id,
		buf: nil,
		len: 0,
	}
}

func (p *PacketBuilder)Reset(id int32)(*PacketBuilder){
	p.id = id
	p.buf = p.buf[:]
	p.len = 0
	return p
}

func (p *PacketBuilder)Id()(int32){
	return p.id
}

func (p *PacketBuilder)Bytes()(buf []byte){
	bf := make([]byte, 5 + p.len)
	n := encodeVarInt(bf[:5], (int32)(p.id))
	n += copy(bf[n:], p.buf[:p.len])
	buf = make([]byte, 5 + n)
	bf = bf[:n]
	n = encodeVarInt(buf[:5], (int32)(n))
	n += copy(buf[n:], bf)
	return buf[:n]
}

func (p *PacketBuilder)WriteTo(w io.Writer)(n int64, err error){
	var n0 int
	var buf [5]byte
	n0 = encodeVarInt(buf[:], (int32)(p.len))
	n0, err = w.Write(buf[:n0])
	n += (int64)(n0)
	if err != nil {
		return
	}
	n0 = encodeVarInt(buf[:], (int32)(p.id))
	n0, err = w.Write(buf[:n0])
	n += (int64)(n0)
	if err != nil {
		return
	}
	n0, err = w.Write(p.buf[:p.len])
	n += (int64)(n0)
	if err != nil {
		return
	}
	return
}

func (p *PacketBuilder)grow(n int){
	l := p.len + n
	if len(p.buf) >= l {
		return
	}
	if cap(p.buf) >= l {
		p.buf = p.buf[:l]
		return
	}
	newl := cap(p.buf) * 2 + n
	buf := make([]byte, newl)
	copy(buf, p.buf)
	p.buf = buf
}

func (p *PacketBuilder)ByteArray(buf []byte)(*PacketBuilder){
	l := len(buf)
	p.grow(l)

	e := p.len + l
	copy(p.buf[p.len:e], buf)
	p.len = e
	return p
}

func (p *PacketBuilder)Bool(v bool)(*PacketBuilder){
	p.grow(1)

	if v {
		p.buf[p.len] = 0x00
	}else{
		p.buf[p.len] = 0x01
	}
	p.len++
	return p
}

func (p *PacketBuilder)Byte(v int8)(*PacketBuilder){
	p.grow(1)

	p.buf[p.len] = (byte)(v)
	p.len++
	return p
}

func (p *PacketBuilder)Short(v int16)(*PacketBuilder){
	p.grow(2)

	encodeUint16(p.buf[p.len:p.len + 2], (uint16)(v))
	p.len += 2
	return p
}

func (p *PacketBuilder)Int(v int32)(*PacketBuilder){
	p.grow(4)

	encodeUint32(p.buf[p.len:p.len + 4], (uint32)(v))
	p.len += 4
	return p
}

func (p *PacketBuilder)Long(v int64)(*PacketBuilder){
	p.grow(8)

	encodeUint64(p.buf[p.len:p.len + 8], (uint64)(v))
	p.len += 8
	return p
}

func (p *PacketBuilder)VarInt(v int32)(*PacketBuilder){
	p.grow(5)

	n := encodeVarInt(p.buf[p.len:p.len + 5], v)
	p.len += n
	return p
}

func (p *PacketBuilder)VarLong(v int64)(*PacketBuilder){
	p.grow(10)

	n := encodeVarLong(p.buf[p.len:p.len + 10], v)
	p.len += n
	return p
}

func (p *PacketBuilder)Float(v float32)(*PacketBuilder){
	p.grow(4)

	encodeFloat32(p.buf[p.len:p.len + 4], v)
	p.len += 4
	return p
}

func (p *PacketBuilder)Double(v float64)(*PacketBuilder){
	p.grow(8)

	encodeFloat64(p.buf[p.len:p.len + 8], v)
	p.len += 8
	return p
}

func (p *PacketBuilder)UUID(v UUID)(*PacketBuilder){
	p.grow(16)

	s := p.len
	p.len += 16
	copy(p.buf[s:p.len], v[:])
	return p
}

func (p *PacketBuilder)String(v string)(*PacketBuilder){
	p.
		VarInt((int32)(len(v))).
		ByteArray(([]byte)(v))
	return p
}

func (p *PacketBuilder)JSON(v any)(*PacketBuilder){
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	p.
		VarInt((int32)(len(buf))).
		ByteArray(buf)
	return p
}

type PacketReader struct{
	id int32
	buf []byte
	off int
}

func ReadPacket(r io.Reader)(p *PacketReader, err error){
	var (
		n int
		size int32
	)
	if n, size, err = readVarInt(r); err != nil {
		return
	}
	p = &PacketReader{
		buf: make([]byte, (int32)(n) + size),
		off: (int)(n),
	}
	encodeVarInt(p.buf[:n], size)
	if _, err = io.ReadFull(r, p.buf[p.off:]); err != nil {
		return nil, err
	}
	var ok bool
	if p.id, ok = p.VarInt(); !ok {
		return nil, io.EOF
	}
	return
}

func (p *PacketReader)Id()(int32){
	return p.id
}

func (p *PacketReader)Bytes()([]byte){
	return p.buf
}

func (p *PacketReader)Offset()(int){
	return p.off
}

func (p *PacketReader)Remain()(int){
	return len(p.buf) - p.off
}

func (p *PacketReader)ByteArray(buf []byte)(ok bool){
	e := p.off + len(buf)
	if e > len(p.buf) {
		return false
	}
	copy(buf, p.buf[p.off:e])
	p.off = e
	return true
}

func (p *PacketReader)Bool()(v bool, ok bool){
	if p.off >= len(p.buf) {
		return false, false
	}
	v = p.buf[p.off] != 0x00
	p.off++
	return v, true
}

func (p *PacketReader)Byte()(v int8, ok bool){
	if p.off >= len(p.buf) {
		return 0, false
	}
	v = (int8)(p.buf[p.off])
	p.off++
	return v, true
}

func (p *PacketReader)Short()(v int16, ok bool){
	e := p.off + 2
	if e > len(p.buf) {
		return 0, false
	}
	v = (int16)(decodeUint16(p.buf[p.off:e]))
	p.off = e
	return v, true
}

func (p *PacketReader)Int()(v int32, ok bool){
	e := p.off + 4
	if e > len(p.buf) {
		return 0, false
	}
	v = (int32)(decodeUint32(p.buf[p.off:e]))
	p.off = e
	return v, true
}

func (p *PacketReader)Long()(v int64, ok bool){
	e := p.off + 8
	if e > len(p.buf) {
		return 0, false
	}
	v = (int64)(decodeUint64(p.buf[p.off:e]))
	p.off = e
	return v, true
}

func (p *PacketReader)Uint8()(v uint8, ok bool){
	var v0 int8
	if v0, ok = p.Byte(); !ok {
		return
	}
	v = (uint8)(v0)
	return
}

func (p *PacketReader)Uint16()(v uint16, ok bool){
	var v0 int16
	if v0, ok = p.Short(); !ok {
		return
	}
	v = (uint16)(v0)
	return
}

func (p *PacketReader)Uint32()(v uint32, ok bool){
	var v0 int32
	if v0, ok = p.Int(); !ok {
		return
	}
	v = (uint32)(v0)
	return
}

func (p *PacketReader)Uint64()(v uint64, ok bool){
	var v0 int64
	if v0, ok = p.Long(); !ok {
		return
	}
	v = (uint64)(v0)
	return
}

func (p *PacketReader)VarInt()(v int32, ok bool){
	var n int
	n, v = decodeVarInt(p.buf[p.off:])
	p.off += n
	return v, true
}

func (p *PacketReader)VarLong()(v int64, ok bool){
	var n int
	n, v = decodeVarLong(p.buf[p.off:])
	p.off += n
	return v, true
}

func (p *PacketReader)Float()(v float32, ok bool){
	e := p.off + 4
	if e > len(p.buf) {
		return 0, false
	}
	v = decodeFloat32(p.buf[p.off:e])
	p.off = e
	return v, true
}

func (p *PacketReader)Double()(v float64, ok bool){
	e := p.off + 8
	if e > len(p.buf) {
		return
	}
	v = decodeFloat64(p.buf[p.off:e])
	p.off = e
	return v, true
}

func (p *PacketReader)UUID()(v UUID, ok bool){
	p.ByteArray(v[:])
	return v, true
}

func (p *PacketReader)String()(v string, ok bool){
	var size int32
	if size, ok = p.VarInt(); !ok {
		return
	}
	var buf = make([]byte, size)
	if ok = p.ByteArray(buf); !ok {
		return
	}
	v = (string)(buf)
	return
}

func (p *PacketReader)JSON(ptr any)(err error){
	var (
		size int32
		ok bool
	)
	if size, ok = p.VarInt(); !ok {
		return io.EOF
	}
	buf := make([]byte, size)
	if ok = p.ByteArray(buf); !ok {
		return io.EOF
	}
	if err = json.Unmarshal(buf, ptr); err != nil {
		return err
	}
	return
}
