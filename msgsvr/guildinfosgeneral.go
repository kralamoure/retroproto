// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildInfosGeneral struct{}

func (m GuildInfosGeneral) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildInfosGeneral
}

func (m GuildInfosGeneral) Serialized() (string, error) {
	return "", nil
}

func (m *GuildInfosGeneral) Deserialize(extra string) error {
	return nil
}
