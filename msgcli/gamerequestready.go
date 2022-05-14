// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameRequestReady struct{}

func (m GameRequestReady) MessageId() retroproto.MsgCliId {
	return retroproto.GameRequestReady
}

func (m GameRequestReady) MessageName() string {
	return "GameRequestReady"
}

func (m GameRequestReady) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameRequestReady) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
