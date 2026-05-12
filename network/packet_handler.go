package network

import (
	"errors"
	"ethene/network/packets/client/handshaking"
	"ethene/network/packets/client/status"
	status2 "ethene/network/packets/server/status"
	"ethene/network/util"
	"fmt"
	"io"
	"log"
	"net"
)

func HandlePacket(payload []byte, session *Connection) error {
	buffer := util.NewNetworkBufferFromBytes(payload)

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
		return handleConfiguringPackets(id, &buffer, session)
	}

	return nil
}

func handleHandshakePackets(id int32, buffer *util.NetworkBuffer, session *Connection) error {
	if id == 0 {
		packet := handshaking.HandshakePacket{}
		err := packet.Unmarshal(*buffer)
		if err != nil {
			return fmt.Errorf("unmarshal packet: %w", err)
		}
		session.state = State(packet.State)
		log.Printf("Handshake received: Version=%d, Server=%s, Port=%d, State=%d\n", packet.Version, packet.ServerName, packet.Port, packet.State)
		return nil
	}

	return errors.New("unknown handshake packet id")
}

func handleStatusPackets(id int32, buffer *util.NetworkBuffer, session *Connection) error {
	if id == 0 {
		response := &status2.ResponseStatus{
			Version: status2.ResponseStatusVersion{
				Name:     "26.1.2",
				Protocol: 775,
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

func handleConfiguringPackets(id int32, buffer *util.NetworkBuffer, session *Connection) error {
	return errors.New("unknown status packet id")
}
