package liter

import (
	"io"
)

type ChunkSection struct { // 1.9.4 ~ 1.11.2
	/*
	 * Number of non-air blocks present in the chunk section. "Non-air" is defined as any fluid and block other than air, cave air, and void air.
	 * The client will keep count of the blocks as they are broken and placed, and,
	 * if the block count reaches 0, the whole chunk section is not rendered, even if it still has blocks.
	 */
	BlockCount Short
	/* Consists of 4096 entries, representing all the blocks in the chunk section */
	BlockStates PalettedContainer
	/* Consists of 64 entries, representing 4x4x4 biome regions in the chunk section */
	Biomes PalettedContainer
}

var _ Packet = (*ChunkSection)(nil)

func (c *ChunkSection) Encode(b *PacketBuilder) {
	b.Short(c.BlockCount)
}

func (c *ChunkSection) DecodeFrom(r *PacketReader) (err error) {
	var ok bool
	if c.BlockCount, ok = r.Short(); !ok {
		return io.EOF
	}
	return
}

type PalettedContainer struct {
	_for int // 1: block; 2: biome
	/*
	 * Determines how many bits are used to encode entries. Note that not all numbers are valid here.
	 * Single valued: when bits per entry is equal to 0
	 * Indirect:
	 *  - For block states with bits per entry <= 4, 4 bits are used to represent a block.
	 *  - For block states and bits per entry between 5 and 8, the given value is used.
	 *  - For biomes the given value is always used, and will be <= 3
	 * Direct: bits per entry values greater than or equal to a threshold (9 for block states, 4 for biomes).
	 */
	BitsPerEntry UByte
	/*
	 * The bits per entry value determines what format is used for the palette.
	 * In most cases, invalid values will be interpreted as a different value when parsed by the notchian client,
	 * meaning that chunk data will be parsed incorrectly if you use an invalid bits per entry.
	 * Servers must make sure that the bits per entry value is correct.
	 */
	Palette any
	/* Compacted list of indices pointing to entry IDs in the Palette. When Bits Per Entry is 0, this array is empty (see Single valued palette) */
	DataArray []Long
}

func (c PalettedContainer) Encode(b *PacketBuilder) {
	b.UByte(c.BitsPerEntry)
}

func (c *PalettedContainer) DecodeFrom(r *PacketReader) (err error) {
	var ok bool
	if c.BitsPerEntry, ok = r.UByte(); !ok {
		return io.EOF
	}
	return
}

type SingleValuedPalette struct {
	Value VarInt
}

type IndirectPalette struct {
	Palette []VarInt
}

type DirectPalette struct{}
