
package liter

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type (
	Bool   = bool
	Byte   = int8
	UByte  = uint8
	Short  = int16
	UShort = uint16
	Int    = int32
	Long   = int64
	Float  = float32
	Double = float64

	VarInt  int32
	VarLong int64

	String     = string
	Identifier = String
	Object     = map[string]any
	Position   struct {
		X int
		Y int
		Z int
	}
	Angle      = UByte
	UUID       = uuid.UUID
	ByteArray  = []byte

	Optional[T any] struct {
		Ok bool
		V T
	}

	Encodable interface {
		Encode(p *PacketBuilder)
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

var _ Packet = (*Position)(nil)

func (p Position)AsLong()(n Long){
	return (((Long)(p.X) & 0x3ffffff) << 38) | (((Long)(p.Z) & 0x3ffffff) << 12) | ((Long)(p.Y) & 0xfff)
}

func (p Position)Encode(b *PacketBuilder){
	b.Long(p.AsLong())
}

func (p *Position)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	var v Long
	if v, ok = r.Long(); !ok {
		return io.EOF
	}
	*p = PositionFromLong(v)
	return
}

var _ Encodable = (*Optional[Int])(nil)

func (o Optional[T])Encode(b *PacketBuilder){
	if o.Ok {
		b.Encode(o.V)
	}
}

func (o Optional[T])Assert()(v T){
	if !o.Ok {
		panic("assert: optional value is not assigned")
	}
	return o.V
}

func (o *Optional[T])Set(v T){
	o.Ok = true
	o.V = v
}

type PacketBuilder struct {
	protocol int
	id VarInt
	buf []byte
	len int
}

func NewPacket(protocol int, id VarInt)(p *PacketBuilder){
	return &PacketBuilder{
		protocol: protocol,
		id: id,
		buf: nil,
		len: 0,
	}
}

func (p *PacketBuilder)Reset(protocol int, id VarInt)(*PacketBuilder){
	p.protocol = protocol
	p.id = id
	p.buf = p.buf[:0]
	p.len = 0
	return p
}

func (p *PacketBuilder)Protocol()(int){
	return p.protocol
}

func (p *PacketBuilder)Id()(VarInt){
	return p.id
}

func (p *PacketBuilder)Len()(int){
	return p.len
}

func (p *PacketBuilder)Data()([]byte){
	return p.buf[:p.len]
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
	n0, err = w.Write(p.Bytes())
	n = (int64)(n0)
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

func (p *PacketBuilder)Byte(v Byte)(*PacketBuilder){
	p.grow(1)

	p.buf[p.len] = (byte)(v)
	p.len++
	return p
}

func (p *PacketBuilder)UByte(v UByte)(*PacketBuilder){
	p.grow(1)

	p.buf[p.len] = (byte)(v)
	p.len++
	return p
}

func (p *PacketBuilder)Short(v Short)(*PacketBuilder){
	p.grow(2)

	encodeUint16(p.buf[p.len:p.len + 2], (uint16)(v))
	p.len += 2
	return p
}

func (p *PacketBuilder)UShort(v UShort)(*PacketBuilder){
	p.grow(2)

	encodeUint16(p.buf[p.len:p.len + 2], (uint16)(v))
	p.len += 2
	return p
}

func (p *PacketBuilder)Int(v Int)(*PacketBuilder){
	p.grow(4)

	encodeUint32(p.buf[p.len:p.len + 4], (uint32)(v))
	p.len += 4
	return p
}

func (p *PacketBuilder)Long(v Long)(*PacketBuilder){
	p.grow(8)

	encodeUint64(p.buf[p.len:p.len + 8], (uint64)(v))
	p.len += 8
	return p
}

func (p *PacketBuilder)VarInt(v VarInt)(*PacketBuilder){
	p.grow(5)

	n := encodeVarInt(p.buf[p.len:p.len + 5], v)
	p.len += n
	return p
}

func (p *PacketBuilder)VarLong(v VarLong)(*PacketBuilder){
	p.grow(10)

	n := encodeVarLong(p.buf[p.len:p.len + 10], v)
	p.len += n
	return p
}

func (p *PacketBuilder)Float(v Float)(*PacketBuilder){
	p.grow(4)

	encodeFloat32(p.buf[p.len:p.len + 4], v)
	p.len += 4
	return p
}

func (p *PacketBuilder)Double(v Double)(*PacketBuilder){
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

func (p *PacketBuilder)String(v String)(*PacketBuilder){
	p.
		VarInt((VarInt)(len(v))).
		ByteArray(([]byte)(v))
	return p
}

func (p *PacketBuilder)JSON(v any)(*PacketBuilder){
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	p.
		VarInt((VarInt)(len(buf))).
		ByteArray(buf)
	return p
}

func (p *PacketBuilder)Chat(v *Chat)(*PacketBuilder){
	buf, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}
	p.
		VarInt((VarInt)(len(buf))).
		ByteArray(buf)
	return p
}

func (p *PacketBuilder)Encode(v any){
	switch w := v.(type) {
	case Bool:
		p.Bool(w)
	case Byte:
		p.Byte(w)
	case UByte:
		p.UByte(w)
	case Short:
		p.Short(w)
	case UShort:
		p.UShort(w)
	case Int:
		p.Int(w)
	case Long:
		p.Long(w)
	case Float:
		p.Float(w)
	case Double:
		p.Double(w)
	case VarInt:
		p.VarInt(w)
	case VarLong:
		p.VarLong(w)
	case ByteArray:
		p.ByteArray(w)
	case String:
		p.String(w)
	case Object:
		p.JSON(w)
	case UUID:
		p.UUID(w)
	case *Chat:
		p.Chat(w)
	case Encodable:
		w.Encode(p)
	default:
		panic("Unknown non-encodable type")
	}
}


type PacketReader struct {
	protocol int
	id VarInt
	buf []byte
	off int
}

func ReadPacket(protocol int, r io.Reader)(p *PacketReader, err error){
	var (
		n int
		size int32
		id int32
	)
	if _, size, err = readVarInt(r); err != nil {
		return
	}
	if n, id, err = readVarInt(r); err != nil {
		return
	}
	size -= (int32)(n)
	p = &PacketReader{
		protocol: protocol,
		id: (VarInt)(id),
		buf: make([]byte, size),
	}
	if _, err = io.ReadFull(r, p.buf); err != nil {
		return nil, err
	}
	return
}

func ReadPacketFromBytes(protocol int, buf []byte)(p *PacketReader){
	return &PacketReader{
		protocol: protocol,
		buf: buf,
	}
}

func (p *PacketReader)Protocol()(int){
	return p.protocol
}

func (p *PacketReader)Id()(VarInt){
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

func (p *PacketReader)ReadAll()(buf []byte){
	buf = make([]byte, len(p.buf) - p.off)
	copy(buf, p.buf[p.off:])
	p.off = len(p.buf)
	return
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

func (p *PacketReader)Byte()(v Byte, ok bool){
	if p.off >= len(p.buf) {
		return 0, false
	}
	v = (Byte)(p.buf[p.off])
	p.off++
	return v, true
}

func (p *PacketReader)UByte()(v UByte, ok bool){
	if p.off >= len(p.buf) {
		return 0, false
	}
	v = (UByte)(p.buf[p.off])
	p.off++
	return v, true
}

func (p *PacketReader)Short()(v Short, ok bool){
	e := p.off + 2
	if e > len(p.buf) {
		return 0, false
	}
	v = (Short)(decodeUint16(p.buf[p.off:e]))
	p.off = e
	return v, true
}

func (p *PacketReader)UShort()(v UShort, ok bool){
	e := p.off + 2
	if e > len(p.buf) {
		return 0, false
	}
	v = (UShort)(decodeUint16(p.buf[p.off:e]))
	p.off = e
	return v, true
}

func (p *PacketReader)Int()(v Int, ok bool){
	e := p.off + 4
	if e > len(p.buf) {
		return 0, false
	}
	v = (Int)(decodeUint32(p.buf[p.off:e]))
	p.off = e
	return v, true
}

func (p *PacketReader)Long()(v Long, ok bool){
	e := p.off + 8
	if e > len(p.buf) {
		return 0, false
	}
	v = (Long)(decodeUint64(p.buf[p.off:e]))
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

func (p *PacketReader)VarInt()(v VarInt, ok bool){
	var n int
	n, v = decodeVarInt(p.buf[p.off:])
	if n < 0 {
		return 0, false
	}
	p.off += n
	return v, true
}

func (p *PacketReader)VarLong()(v VarLong, ok bool){
	var n int
	n, v = decodeVarLong(p.buf[p.off:])
	if n < 0 {
		return 0, false
	}
	p.off += n
	return v, true
}

func (p *PacketReader)Float()(v Float, ok bool){
	e := p.off + 4
	if e > len(p.buf) {
		return 0, false
	}
	v = decodeFloat32(p.buf[p.off:e])
	p.off = e
	return v, true
}

func (p *PacketReader)Double()(v Double, ok bool){
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

func (p *PacketReader)String()(v String, ok bool){
	var size VarInt
	if size, ok = p.VarInt(); !ok {
		return
	}
	buf := make([]byte, size)
	if ok = p.ByteArray(buf); !ok {
		return
	}
	v = (String)(buf)
	return
}

func (p *PacketReader)JSON(ptr any)(err error){
	var (
		size VarInt
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
		return
	}
	return
}

func (p *PacketReader)Chat()(c *Chat, err error){
	var (
		size VarInt
		ok bool
	)
	if size, ok = p.VarInt(); !ok {
		return nil, io.EOF
	}
	buf := make([]byte, size)
	if ok = p.ByteArray(buf); !ok {
		return nil, io.EOF
	}
	c = new(Chat)
	if err = c.UnmarshalJSON(buf); err != nil {
		return nil, err
	}
	return
}
