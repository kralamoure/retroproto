package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameMapLoaded struct{}

func (m GameMapLoaded) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameMapLoaded
}

func (m GameMapLoaded) Serialized() (string, error) {
	return "", nil
}

func (m *GameMapLoaded) Deserialize(extra string) error {
	return nil
}
