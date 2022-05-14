// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameLeave struct{}

func (m GameLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.GameLeave
}

func (m GameLeave) MessageName() string {
	return "GameLeave"
}

func (m GameLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
