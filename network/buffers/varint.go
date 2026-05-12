package buffers

import (
	"fmt"
	"io"
)

func ReadVarInt(r io.ByteReader) (int32, error) {
	var result int32
	var shift uint = 0
	var b byte

	for {
		var err error
		b, err = r.ReadByte()
		if err != nil {
			return 0, err
		}

		result |= int32(b&0x7F) << shift
		shift += 7

		if b&0x80 == 0 {
			break
		}

		if shift >= 35 {
			return 0, fmt.Errorf("VarInt too big")
		}
	}

	return result, nil
}

func WriteVarInt(w io.ByteWriter, value int32) error {
	var b byte

	for {
		b = byte(value & 0x7F)
		value >>= 7

		if value != 0 {
			b |= 0x80
		}

		if err := w.WriteByte(b); err != nil {
			return err
		}

		if value == 0 {
			break
		}
	}

	return nil
}

func VarIntLength(value int32) int {
	length := 0
	for {
		length++
		value >>= 7
		if value == 0 {
			break
		}
	}
	return length
}
