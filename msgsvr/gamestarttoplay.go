// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameStartToPlay struct{}

func (m GameStartToPlay) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameStartToPlay
}

func (m GameStartToPlay) Serialized() (string, error) {
	return "", nil
}

func (m *GameStartToPlay) Deserialize(extra string) error {
	return nil
}
