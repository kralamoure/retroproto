package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameCreateError struct{}

func (m GameCreateError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameCreateError
}

func (m GameCreateError) Serialized() (string, error) {
	return "", nil
}

func (m *GameCreateError) Deserialize(extra string) error {
	return nil
}
