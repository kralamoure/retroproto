// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameEnd struct{}

func (m GameEnd) MessageId() retroproto.MsgSvrId {
	return retroproto.GameEnd
}

func (m GameEnd) MessageName() string {
	return "GameEnd"
}

func (m GameEnd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameEnd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
