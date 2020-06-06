// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GameSetPlayerPosition struct{}

func (m GameSetPlayerPosition) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameSetPlayerPosition
}

func (m GameSetPlayerPosition) Serialized() (string, error) {
	return "", nil
}

func (m *GameSetPlayerPosition) Deserialize(extra string) error {
	return nil
}
