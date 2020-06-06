// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameFrameObject2 struct{}

func (m GameFrameObject2) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameFrameObject2
}

func (m GameFrameObject2) Serialized() (string, error) {
	return "", nil
}

func (m *GameFrameObject2) Deserialize(extra string) error {
	return nil
}
