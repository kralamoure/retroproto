// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameLeave struct{}

func (m GameLeave) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameLeave
}

func (m GameLeave) Serialized() (string, error) {
	return "", nil
}

func (m *GameLeave) Deserialize(extra string) error {
	return nil
}
