package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type ExchangeRequest struct {
	Type int
	Id   int
	Cell int
}

func (m ExchangeRequest) ProtocolId() retroproto.MsgCliId {
	return retroproto.ExchangeRequest
}

func (m ExchangeRequest) Serialized() (string, error) {
	var id string
	if m.Id != 0 {
		id = fmt.Sprintf("%d", m.Id)
	}

	var cell string
	if m.Cell != 0 {
		id = fmt.Sprintf("|%d", m.Cell)
	}

	return fmt.Sprintf("%d|%s%s", m.Type, id, cell), nil
}

func (m *ExchangeRequest) Deserialize(extra string) error {

	sli := strings.Split(extra, "|")

	if len(sli) < 2 {
		return retroproto.ErrInvalidMsg
	}

	requestType, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Type = int(requestType)

	if sli[1] != "" {
		id, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}
		m.Id = int(id)
	}

	if len(sli) >= 3 && sli[2] != "" {
		cell, err := strconv.ParseInt(sli[2], 10, 32)
		if err != nil {
			return err
		}
		m.Cell = int(cell)
	}

	return nil
}
