// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameFightOption struct{}

func (m GameFightOption) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameFightOption
}

func (m GameFightOption) Serialized() (string, error) {
	return "", nil
}

func (m *GameFightOption) Deserialize(extra string) error {
	return nil
}
