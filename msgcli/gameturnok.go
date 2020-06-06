// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GameTurnOk struct{}

func (m GameTurnOk) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameTurnOk
}

func (m GameTurnOk) Serialized() (string, error) {
	return "", nil
}

func (m *GameTurnOk) Deserialize(extra string) error {
	return nil
}
