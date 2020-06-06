// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameFlag struct{}

func (m GameFlag) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameFlag
}

func (m GameFlag) Serialized() (string, error) {
	return "", nil
}

func (m *GameFlag) Deserialize(extra string) error {
	return nil
}
