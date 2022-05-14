// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameReady struct{}

func (m GameReady) MessageId() retroproto.MsgSvrId {
	return retroproto.GameReady
}

func (m GameReady) MessageName() string {
	return "GameReady"
}

func (m GameReady) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameReady) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
