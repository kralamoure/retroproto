// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GameSetFlag struct{}

func (m GameSetFlag) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameSetFlag
}

func (m GameSetFlag) Serialized() (string, error) {
	return "", nil
}

func (m *GameSetFlag) Deserialize(extra string) error {
	return nil
}
