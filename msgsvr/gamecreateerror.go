package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameCreateError struct{}

func (m GameCreateError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameCreateError
}

func (m GameCreateError) Serialized() (string, error) {
	return "", nil
}

func (m *GameCreateError) Deserialize(extra string) error {
	return nil
}
