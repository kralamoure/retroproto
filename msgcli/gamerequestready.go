// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameRequestReady struct{}

func (m GameRequestReady) ProtocolId() retroproto.MsgCliId {
	return retroproto.GameRequestReady
}

func (m GameRequestReady) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameRequestReady) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
