package buffers

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/google/uuid"
)

type NetworkBuffer interface {
	WriteByte(value byte) error
	ReadByte() (byte, error)
	WriteBytes(value []byte)
	ReadBytes() []byte
	WriteShort(value int16)
	ReadShort() int16
	WriteInt(value int32)
	ReadInt() int32
	WriteLong(value int64)
	ReadLong() int64
	WriteFloat(value float32)
	ReadFloat() float32
	WriteDouble(value float64)
	ReadDouble() float64
	WriteString(value string)
	ReadString() string
	WriteVarInt(value int32)
	ReadVarInt() (int32, error)
	WriteBool(value bool)
	ReadBool() bool
	ReadUUID() uuid.UUID
	Bytes() []byte
}

type NetworkBufferImpl struct {
	buffer *bytes.Buffer
}

func (b *NetworkBufferImpl) ReadVarInt() (int32, error) {
	return ReadVarInt(b)
}

func NewNetworkBufferFromBytes(data []byte) NetworkBuffer {
	return &NetworkBufferImpl{
		buffer: bytes.NewBuffer(data),
	}
}

func (b *NetworkBufferImpl) WriteByte(value byte) error {
	err := b.buffer.WriteByte(value)
	if err != nil {
		fmt.Println("Error writing byte:", err)
		return err
	}
	return nil
}

func (b *NetworkBufferImpl) ReadByte() (byte, error) {
	b2, err := b.buffer.ReadByte()
	if err != nil {
		fmt.Println("Error reading byte:", err)
		return 0, err
	}
	return b2, nil
}

func (b *NetworkBufferImpl) WriteBytes(value []byte) {
	length := len(value)
	varIntLength := VarIntLength(int32(length))
	totalLength := length + varIntLength

	b.WriteVarInt(int32(totalLength))
	b.WriteBytes(value)

	_, err := b.buffer.Write(value)
	if err != nil {
		fmt.Println("Error writing bytes:", err)
	}
}

func (b *NetworkBufferImpl) ReadBytes() []byte {
	return make([]byte, 0) // TODO
}

func (b *NetworkBufferImpl) WriteShort(value int16) {
	err := binary.Write(b.buffer, binary.BigEndian, value)
	if err != nil {
		fmt.Println("Error writing int16:", err)
	}
}

func (b *NetworkBufferImpl) ReadShort() int16 {
	var value int16
	err := binary.Read(b.buffer, binary.BigEndian, &value)
	if err != nil {
		fmt.Println("Error reading int16:", err)
	}
	return value
}

func (b *NetworkBufferImpl) WriteInt(value int32) {
	err := binary.Write(b.buffer, binary.BigEndian, value)
	if err != nil {
		fmt.Println("Error writing int32:", err)
	}
}

func (b *NetworkBufferImpl) ReadInt() int32 {
	var value int32
	err := binary.Read(b.buffer, binary.BigEndian, &value)
	if err != nil {
		fmt.Println("Error reading int32:", err)
	}
	return value
}

func (b *NetworkBufferImpl) WriteLong(value int64) {
	err := binary.Write(b.buffer, binary.BigEndian, value)
	if err != nil {
		fmt.Println("Error writing int64:", err)
	}
}

func (b *NetworkBufferImpl) ReadLong() int64 {
	var value int64
	err := binary.Read(b.buffer, binary.BigEndian, &value)
	if err != nil {
		fmt.Println("Error reading int64:", err)
	}
	return value
}

func (b *NetworkBufferImpl) WriteFloat(value float32) {
	err := binary.Write(b.buffer, binary.BigEndian, value)
	if err != nil {
		fmt.Println("Error writing float32:", err)
	}
}

func (b *NetworkBufferImpl) ReadFloat() float32 {
	var value float32
	err := binary.Read(b.buffer, binary.BigEndian, &value)
	if err != nil {
		fmt.Println("Error reading float32:", err)
	}
	return value
}

func (b *NetworkBufferImpl) WriteDouble(value float64) {
	err := binary.Write(b.buffer, binary.BigEndian, value)
	if err != nil {
		fmt.Println("Error writing float64:", err)
	}
}

func (b *NetworkBufferImpl) ReadDouble() float64 {
	var value float64
	err := binary.Read(b.buffer, binary.BigEndian, &value)
	if err != nil {
		fmt.Println("Error reading float64:", err)
	}
	return value
}

func (b *NetworkBufferImpl) WriteString(value string) {
	length := int32(len(value))
	b.WriteVarInt(length)
	_, err := b.buffer.WriteString(value)
	if err != nil {
		fmt.Println("Error writing string:", err)
	}
}

func (b *NetworkBufferImpl) ReadString() string {
	length, err := b.ReadVarInt()
	if err != nil {
		fmt.Println("Error reading string:", err)
	}
	value := make([]byte, length)
	_, err = b.buffer.Read(value)
	if err != nil {
		fmt.Println("Error reading string:", err)
	}
	return string(value)
}

func (b *NetworkBufferImpl) WriteBool(value bool) {
	err := binary.Write(b.buffer, binary.BigEndian, value)
	if err != nil {
		fmt.Println("Error writing bool:", err)
	}
}

func (b *NetworkBufferImpl) ReadBool() bool {
	var value bool
	err := binary.Read(b.buffer, binary.BigEndian, &value)
	if err != nil {
		fmt.Println("Error reading bool:", err)
	}
	return value
}

func (b *NetworkBufferImpl) WriteVarInt(value int32) {
	err := WriteVarInt(b.buffer, value)
	if err != nil {
		fmt.Println("Error writing varint:", err)
	}
}

func (b *NetworkBufferImpl) ReadUUID() uuid.UUID {
	var value uuid.UUID
	err := binary.Read(b.buffer, binary.BigEndian, &value)
	if err != nil {
		fmt.Println("Error reading uuid:", err)
	}

	return value
}

func (b *NetworkBufferImpl) Bytes() []byte {
	return b.buffer.Bytes()
}
