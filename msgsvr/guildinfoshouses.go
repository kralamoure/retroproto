// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildInfosHouses struct{}

func (m GuildInfosHouses) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildInfosHouses
}

func (m GuildInfosHouses) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildInfosHouses) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
