
package main

import (
	"strings"
)

func wikiTypeAsGoType(wikiTyp string)(typ string){
	wikiTyp = strings.ToLower(wikiTyp)
	if wikiTyp == "array of (byte, byte, byte)" {
		return "[][3]Byte"
	}
	wikiTyp, _ = split(wikiTyp, ',')
	wikiTyp, _ = split(wikiTyp, '(')
	wikiTyp = strings.TrimSpace(wikiTyp)
	wikiTyp = strings.TrimSuffix(wikiTyp, " enum")
	if cutted, ok := strings.CutPrefix(wikiTyp, "optional "); ok {
		return "Optional[" + wikiTypeAsGoType(cutted) + "]"
	}
	if cutted, ok := strings.CutPrefix(wikiTyp, "array of "); ok {
		return "[]" + wikiTypeAsGoType(cutted)
	}
	if cutted, ok := strings.CutSuffix(wikiTyp, " array"); ok {
		if cutted == "byte" {
			return "ByteArray"
		}
		return "[]" + wikiTypeAsGoType(cutted)
	}
	switch wikiTyp {
	case "boolean":
		return "Bool"
	case "byte":
		return "Byte"
	case "unsigned byte":
		return "UByte"
	case "short":
		return "Short"
	case "unsigned short":
		return "UShort"
	case "int", "integer":
		return "Int"
	case "long", "instant":
		return "Long"
	case "float":
		return "Float"
	case "double":
		return "Double"
	case "varint":
		return "VarInt"
	case "varlong":
		return "VarLong"
	case "metadata", "entity metadata":
		return "*data.EntityMetadata"
	case "slot":
		return "*data.Slot"
	case "nbt", "nbt tag":
		return "nbt.NBT"
	case "nbt tag compound":
		return "*nbt.NBTCompound"
	case "string", "identifier":
		return "String"
	case "chat", "component":
		return "Object"
	case "position":
		return "Position"
	case "angle":
		return "Angle"
	case "uuid":
		return "UUID"
	case "varies":
		return "any"
	case "bitset":
		return "BitSet"
	case "node":
		return "*CommandNode"
	case "chunk", "chunk section":
		return "*ChunkSection"
	}
	if cutted, ok := strings.CutSuffix(wikiTyp, "s"); ok {
		return "[]" + wikiTypeAsGoType(cutted)
	}
	loger.Warnf("Unknown type %q", wikiTyp)
	return "Unknown_" + strings.ReplaceAll(wikiTyp, " ", "_")
}
