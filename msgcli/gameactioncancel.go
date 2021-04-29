package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type GameActionCancel struct {
	Id     int
	Params string
}

func (m GameActionCancel) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameActionCancel
}

func (m GameActionCancel) Serialized() (string, error) {
	return fmt.Sprintf("%d|%s", m.Id, m.Params), nil
}

func (m *GameActionCancel) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")

	if sli[0] == "" {
		return d1proto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	if len(sli) >= 2 {
		m.Params = sli[1]
	}

	return nil
}
