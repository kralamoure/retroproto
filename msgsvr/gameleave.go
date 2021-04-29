// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameLeave struct{}

func (m GameLeave) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameLeave
}

func (m GameLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
