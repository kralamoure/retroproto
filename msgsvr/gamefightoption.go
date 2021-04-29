// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameFightOption struct{}

func (m GameFightOption) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameFightOption
}

func (m GameFightOption) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameFightOption) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
