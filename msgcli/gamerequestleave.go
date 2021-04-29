// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GameRequestLeave struct{}

func (m GameRequestLeave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GameRequestLeave
}

func (m GameRequestLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameRequestLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
