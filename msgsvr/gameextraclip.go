// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameExtraClip struct{}

func (m GameExtraClip) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameExtraClip
}

func (m GameExtraClip) Serialized() (string, error) {
	return "", nil
}

func (m *GameExtraClip) Deserialize(extra string) error {
	return nil
}
