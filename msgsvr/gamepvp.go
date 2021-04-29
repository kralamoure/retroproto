// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GamePVP struct{}

func (m GamePVP) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GamePVP
}

func (m GamePVP) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GamePVP) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
