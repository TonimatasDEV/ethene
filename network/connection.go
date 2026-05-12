package network

import (
	"bufio"
	"ethene/network/buffers"
	"ethene/network/packets"
	"fmt"
	"io"
	"log"
	"net"
)

type State int

const (
	Handshake State = iota
	Status
	Login
	Configuration
	Play
)

type Connection struct {
	conn  net.Conn
	state State
	r     *bufio.Reader
	w     *bufio.Writer
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		conn: conn,
		r:    bufio.NewReader(conn),
		w:    bufio.NewWriter(conn),
	}
}

func HandleConnection(session *Connection) error {
	defer session.conn.Close()
	session.state = Handshake

	for {
		length, err := buffers.ReadVarInt(session.r)
		if err != nil {
			if err != io.EOF {
				log.Println("Handshake failed:", err)
				return err
			}

			return nil
		}

		payload := make([]byte, length)
		_, err = io.ReadFull(session.r, payload)
		if err != nil {
			return fmt.Errorf("error reading packet payload: %w", err)
		}

		err = HandlePacket(payload, session)
		if err != nil {
			log.Println("Failed processing packet:", err)
			return err
		}
	}
}

func (conn *Connection) SendPacket(packet packets.ServerPacket) error {
	payloadBuffer := buffers.NewNetworkBufferFromBytes(make([]byte, 0))
	packet.Marshal(payloadBuffer)
	payload := payloadBuffer.Bytes()

	idLength := buffers.VarIntLength(packet.Id())
	totalLength := int32(idLength + len(payload))

	buffer := buffers.NewNetworkBufferFromBytes(make([]byte, 0))
	buffer.WriteVarInt(totalLength)
	buffer.WriteVarInt(packet.Id())
	buffer.WriteBytes(payload)

	data := buffer.Bytes()

	if _, err := conn.w.Write(data); err != nil {
		return err
	}
	return conn.w.Flush()
}
