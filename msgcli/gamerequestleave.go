// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GameRequestLeave struct{}

func (m GameRequestLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameRequestLeave
}

func (m GameRequestLeave) Serialized() (string, error) {
	return "", nil
}

func (m *GameRequestLeave) Deserialize(extra string) error {
	return nil
}
