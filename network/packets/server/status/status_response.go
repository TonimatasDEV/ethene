package status

import (
	"encoding/json"
	"ethene/network/buffers"
	"log"
)

type ResponseStatusVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type ResponseStatusPlayersSample struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type ResponseStatusPlayers struct {
	Max    int                           `json:"max"`
	Online int                           `json:"online"`
	Sample []ResponseStatusPlayersSample `json:"sample"`
}

type ResponseStatusDescription struct {
	Text string `json:"text"`
}

type ResponseStatus struct {
	Version            ResponseStatusVersion     `json:"version"`
	Players            ResponseStatusPlayers     `json:"players"`
	Description        ResponseStatusDescription `json:"description"`
	Favicon            string                    `json:"favicon"`
	EnforcesSecureChat bool                      `json:"enforcesSecureChat"`
}

func (p *ResponseStatus) Marshal(buffer buffers.NetworkBuffer) {
	data, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
		return
	}
	buffer.WriteString(string(data))
}

func (p *ResponseStatus) Id() int32 {
	return 0
}
