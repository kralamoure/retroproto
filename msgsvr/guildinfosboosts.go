// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildInfosBoosts struct{}

func (m GuildInfosBoosts) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildInfosBoosts
}

func (m GuildInfosBoosts) Serialized() (string, error) {
	return "", nil
}

func (m *GuildInfosBoosts) Deserialize(extra string) error {
	return nil
}
