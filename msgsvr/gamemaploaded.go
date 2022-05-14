package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameMapLoaded struct{}

func (m GameMapLoaded) MessageId() retroproto.MsgSvrId {
	return retroproto.GameMapLoaded
}

func (m GameMapLoaded) MessageName() string {
	return "GameMapLoaded"
}

func (m GameMapLoaded) Serialized() (string, error) {
	return "", nil
}

func (m *GameMapLoaded) Deserialize(extra string) error {
	return nil
}
