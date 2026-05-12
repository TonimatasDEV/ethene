package network

import (
	"errors"
	"ethene/network/buffers"
	"ethene/network/packets/client/handshaking"
	"ethene/network/packets/client/status"
	status2 "ethene/network/packets/server/status"
	"ethene/util"
	"fmt"
	"io"
	"net"
)

func HandlePacket(payload []byte, session *Connection) error {
	buffer := buffers.NewNetworkBufferFromBytes(payload)

	id, err := buffer.ReadVarInt()
	if err != nil {
		if err == io.EOF || errors.Is(err, net.ErrClosed) {
			fmt.Println("Connection closed")
			return nil
		}
		return fmt.Errorf("reading packet ID: %w", err)
	}

	if session.state == Handshake {
		return handleHandshakePackets(id, &buffer, session)
	} else if session.state == Status {
		return handleStatusPackets(id, &buffer, session)
	} else if session.state == Login {
		return handleLoginPackets(id, &buffer, session)
	}

	return nil
}

func handleHandshakePackets(id int32, buffer *buffers.NetworkBuffer, session *Connection) error {
	if id == 0 {
		packet := handshaking.HandshakePacket{}
		err := packet.Unmarshal(*buffer)
		if err != nil {
			return fmt.Errorf("unmarshal packet: %w", err)
		}
		session.state = State(packet.State)
		return nil
	}

	return errors.New("unknown handshake packet id")
}

func handleStatusPackets(id int32, buffer *buffers.NetworkBuffer, session *Connection) error {
	if id == 0 {
		response := &status2.ResponseStatus{
			Version: status2.ResponseStatusVersion{
				Name:     util.ProtocolName,
				Protocol: util.ProtocolVersion,
			},
			Players: status2.ResponseStatusPlayers{
				Max:    100,
				Online: 0,
				Sample: make([]status2.ResponseStatusPlayersSample, 0),
			},
			Description: status2.ResponseStatusDescription{
				Text: "Hello world!",
			},
			Favicon: "",
		}

		if err := session.SendPacket(response); err != nil {
			return fmt.Errorf("error sending packet: %w", err)
		}

		return nil
	} else if id == 1 {
		ping := status.PingRequest{}

		if err := ping.Unmarshal(*buffer); err != nil {
			return fmt.Errorf("unmarshal packet: %w", err)
		}

		pong := status2.PongResponse{Timestamp: ping.Timestamp}

		if err := session.SendPacket(pong); err != nil {
			return fmt.Errorf("error sending packet: %w", err)
		}

		return nil
	}

	return errors.New("unknown status packet id")
}

func handleLoginPackets(id int32, buffer *buffers.NetworkBuffer, session *Connection) error {
	return errors.New("unknown login packet id")
}
