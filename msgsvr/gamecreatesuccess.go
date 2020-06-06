package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type GameCreateSuccess struct {
	Type int
}

func (m GameCreateSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameCreateSuccess
}

func (m GameCreateSuccess) Serialized() (string, error) {
	return fmt.Sprintf("|%d", m.Type), nil
}

func (m *GameCreateSuccess) Deserialize(extra string) error {
	extra = strings.TrimPrefix(extra, "|")
	if len(extra) < 1 {
		return d1proto.ErrInvalidMsg
	}

	typeValue, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Type = int(typeValue)

	return nil
}
