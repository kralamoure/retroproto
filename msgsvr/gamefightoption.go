// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameFightOption struct{}

func (m GameFightOption) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameFightOption
}

func (m GameFightOption) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameFightOption) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
