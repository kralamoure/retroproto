// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameFrameObjectExternal struct{}

func (m GameFrameObjectExternal) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameFrameObjectExternal
}

func (m GameFrameObjectExternal) Serialized() (string, error) {
	return "", nil
}

func (m *GameFrameObjectExternal) Deserialize(extra string) error {
	return nil
}
