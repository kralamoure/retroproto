// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameRequestLeave struct{}

func (m GameRequestLeave) ProtocolId() retroproto.MsgCliId {
	return retroproto.GameRequestLeave
}

func (m GameRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
