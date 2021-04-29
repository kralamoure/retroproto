package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameMapLoaded struct{}

func (m GameMapLoaded) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameMapLoaded
}

func (m GameMapLoaded) Serialized() (string, error) {
	return "", nil
}

func (m *GameMapLoaded) Deserialize(extra string) error {
	return nil
}
