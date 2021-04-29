// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GameFreeMySoul struct{}

func (m GameFreeMySoul) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GameFreeMySoul
}

func (m GameFreeMySoul) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameFreeMySoul) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
