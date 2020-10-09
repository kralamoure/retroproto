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
	return "", d1proto.ErrNotImplemented
}

func (m *GameRequestLeave) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
