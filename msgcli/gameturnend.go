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
	return "", nil
}

func (m *GameTurnEnd) Deserialize(extra string) error {
	return nil
}
