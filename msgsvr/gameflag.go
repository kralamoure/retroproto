// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameFlag struct{}

func (m GameFlag) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GameFlag
}

func (m GameFlag) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFlag) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
