// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameFreeMySoul struct{}

func (m GameFreeMySoul) ProtocolId() retroproto.MsgCliId {
	return retroproto.GameFreeMySoul
}

func (m GameFreeMySoul) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFreeMySoul) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
