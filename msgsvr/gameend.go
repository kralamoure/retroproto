// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameEnd struct{}

func (m GameEnd) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameEnd
}

func (m GameEnd) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameEnd) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
