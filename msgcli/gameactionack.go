// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GameActionAck struct{}

func (m GameActionAck) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameActionAck
}

func (m GameActionAck) Serialized() (string, error) {
	return "", nil
}

func (m *GameActionAck) Deserialize(extra string) error {
	return nil
}
