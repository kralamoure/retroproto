// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GameTurnEnd struct{}

func (m GameTurnEnd) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameTurnEnd
}

func (m GameTurnEnd) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameTurnEnd) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
