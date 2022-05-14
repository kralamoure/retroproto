// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameFreeMySoul struct{}

func (m GameFreeMySoul) MessageId() retroproto.MsgCliId {
	return retroproto.GameFreeMySoul
}

func (m GameFreeMySoul) MessageName() string {
	return "GameFreeMySoul"
}

func (m GameFreeMySoul) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFreeMySoul) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
