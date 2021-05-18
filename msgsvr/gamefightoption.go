// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameFightOption struct{}

func (m GameFightOption) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GameFightOption
}

func (m GameFightOption) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFightOption) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
