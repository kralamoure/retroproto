package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameCreateError struct{}

func (m GameCreateError) MessageId() retroproto.MsgSvrId {
	return retroproto.GameCreateError
}

func (m GameCreateError) MessageName() string {
	return "GameCreateError"
}

func (m GameCreateError) Serialized() (string, error) {
	return "", nil
}

func (m *GameCreateError) Deserialize(extra string) error {
	return nil
}
