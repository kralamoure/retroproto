// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameEnd struct{}

func (m GameEnd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameEnd
}

func (m GameEnd) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameEnd) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
