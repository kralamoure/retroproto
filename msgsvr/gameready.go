// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameReady struct{}

func (m GameReady) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameReady
}

func (m GameReady) Serialized() (string, error) {
	return "", nil
}

func (m *GameReady) Deserialize(extra string) error {
	return nil
}
