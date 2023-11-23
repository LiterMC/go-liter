package data

import (
	"fmt"
	"io"
	"sync"

	. "github.com/kmcsr/go-liter"
)

type EntityPropEncoder interface {
	Type() VarInt
	Encode(b *PacketBuilder, value any) (err error)
	Decode(r *PacketReader) (value any, err error)
}

var (
	entityEncoderMux sync.RWMutex
	entityEncoders   = make(map[VarInt]EntityPropEncoder, 30)
)

func RegisterEntityEncoder(p EntityPropEncoder) (ok bool) {
	typ := p.Type()

	entityEncoderMux.Lock()
	defer entityEncoderMux.Unlock()
	if _, ok := entityEncoders[typ]; ok {
		return false
	}
	entityEncoders[typ] = p
	return true
}

func UnregisterEntityEncoder(p EntityPropEncoder) (ok bool) {
	typ := p.Type()

	entityEncoderMux.Lock()
	defer entityEncoderMux.Unlock()
	old, ok := entityEncoders[typ]
	if !ok || old != p {
		return false
	}
	delete(entityEncoders, typ)
	return true
}

type EntityMetadata struct {
	/* Unique index key determining the meaning of the following value, see the table below.
	 * If this is 0xff then the it is the end of the Entity Metadata array and no more is read.
	 */
	Index UByte
	/* Only if Index is not 0xff; the type of the index, see the table below */
	Type Optional[VarInt]
	/* Only if Index is not 0xff: the value of the metadata field, see the table below */
	Value any
}

var _ Packet = (*EntityMetadata)(nil)

func (c EntityMetadata) Encode(b *PacketBuilder) {
	b.UByte(c.Index)
	if c.Index != 0xff {
		typ := c.Type.Assert()
		b.VarInt(typ)
		entityEncoderMux.RLock()
		encoder, ok := entityEncoders[typ]
		entityEncoderMux.RUnlock()
		if !ok {
			panic(fmt.Errorf("liter/data.EntityMetadata: Entity metadata encoder of type %d is not exists", typ))
		}
		if err := encoder.Encode(b, c.Value); err != nil {
			panic(err)
		}
	}
}

func (c *EntityMetadata) DecodeFrom(r *PacketReader) (err error) {
	var ok bool
	if c.Index, ok = r.UByte(); !ok {
		return io.EOF
	}
	if c.Type.Ok = c.Index != 0xff; c.Type.Ok {
		if c.Type.V, ok = r.VarInt(); !ok {
			return io.EOF
		}
		entityEncoderMux.RLock()
		encoder, ok := entityEncoders[c.Type.V]
		entityEncoderMux.RUnlock()
		if !ok {
			return fmt.Errorf("liter/data.EntityMetadata: Entity metadata encoder of type %d is not exists", c.Type.V)
		}
		if c.Value, err = encoder.Decode(r); err != nil {
			return
		}
	}
	return
}
