
package nbt

import (
	"sync"
)

var utf8mBufPool = sync.Pool{
	New: func()(any){
		return new([]byte)
	},
}

// Encode a UTF8 string to Java modified UTF8
func EncodeUtf8m(str string)(out []byte){
	out = make([]byte, 0, len(str) * 4 / 3)
	for i := 0; i < len(str); i++ {
		b1 := str[i]
		// NULL bytes and bytes starting with 11110xxx are special
		if (b1 & 0x80) == 0 {
			if b1 == 0 {
				out = append(out, 0xC0)
				out = append(out, 0x80)
			}else{
				// Single byte
				out = append(out, b1)
			}
		}else if (b1 & 0xE0) == 0xC0 { // 2bytes encoding
			out = append(out, b1)
			i++
			out = append(out, str[i])
		}else if (b1 & 0xF0) == 0xE0 { // 3bytes encoding
			out = append(out, b1)
			i++
			out = append(out, str[i])
			i++
			out = append(out, str[i])
		}else if (b1 & 0xF8) == 0xF0 { // 4bytes encoding
			// Beginning of 4byte encoding, turn into 2 3byte encodings
			// Bits in: 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
			i++
			b2 := str[i]
			i++
			b3 := str[i]
			i++
			b4 := str[i]

			// Reconstruct full 21bit value
			u21 := (uint32)(b1 & 0x07) << 18
			u21 += (uint32)(b2 & 0x3F) << 12
			u21 += (uint32)(b3 & 0x3F) << 6
			u21 += (uint32)(b4 & 0x3F)

			// Bits out: 11101101 1010xxxx 10xxxxxx
			out = append(out, 0xED)
			out = append(out, (byte)(0xA0 + (((u21 >> 16) - 1) & 0x0F)))
			out = append(out, (byte)(0x80 + ((u21 >> 10) & 0x3F)))

			// Bits out: 11101101 1011xxxx 10xxxxxx
			out = append(out, 0xED)
			out = append(out, (byte)(0xB0 + ((u21 >> 6) & 0x0F)))
			out = append(out, b4)
		}
	}
	return
}

// Decode Java modified UTF8 to normal UTF8 string
func DecodeUtf8m(str []byte)(string){
	outP := utf8mBufPool.Get().(*[]byte)
	defer utf8mBufPool.Put(outP)
	if *outP == nil {
		*outP = make([]byte, 0, len(str))
	}
	out := (*outP)[:0]
	leng := len(str)
	for i := 0; i < leng; i++ {
		var b1, b2, b3 byte
		b1 = str[i]
		if (b1 & 0x80) == 0 { // 1byte encoding
			out = append(out, b1)
		}else if (b1 & 0xE0) == 0xC0 { // 2bytes encoding
			i++
			b2 = str[i]
			if b1 != 0xC0 || b2 != 0x80 {
				out = append(out, b1)
				out = append(out, b2)
			}else{
				out = append(out, 0x00)
			}
		}else if (b1 & 0xF0) == 0xE0 { // 3bytes encoding
			i++
			b2 = str[i]
			i++
			b3 = str[i]
			if i + 3 < leng && b1 == 0xED && (b2 & 0xF0) == 0xA0 {
				// See if this is a pair of 3byte encodings
				b4 := str[i + 1]
				b5 := str[i + 2]
				b6 := str[i + 3]
				if b4 == 0xED && (b5 & 0xF0) == 0xB0 {
					// Bits in: 11101101 1010xxxx 10xxxxxx
					// Bits in: 11101101 1011xxxx 10xxxxxx
					i += 3
					// Reconstruct 21 bit code
					u21 := ((uint32)(b2 & 0x0F) + 1) << 16
					u21 += (uint32)(b3 & 0x3F) << 10
					u21 += (uint32)(b5 & 0x0F) << 6
					u21 += (uint32)(b6 & 0x3F)
					// Bits out: 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
					// Convert to 4byte encoding
					out = append(out, 0xF0 + (byte)((u21 >> 18) & 0x07))
					out = append(out, 0x80 + (byte)((u21 >> 12) & 0x3F))
					out = append(out, 0x80 + (byte)((u21 >> 6) & 0x3F))
					out = append(out, 0x80 + (byte)(u21 & 0x3F))
					continue
				}
			}
			out = append(out, b1)
			out = append(out, b2)
			out = append(out, b3)
		}
	}
	return (string)(out)
}
