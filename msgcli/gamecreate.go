package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type GameCreate struct {
	Type int
}

func (m GameCreate) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameCreate
}

func (m GameCreate) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Type), nil
}

func (m *GameCreate) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")

	typeValue, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Type = int(typeValue)

	return nil
}
