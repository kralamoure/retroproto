// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameFlag struct{}

func (m GameFlag) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameFlag
}

func (m GameFlag) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameFlag) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
