// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildInfosHouses struct{}

func (m GuildInfosHouses) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GuildInfosHouses
}

func (m GuildInfosHouses) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosHouses) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
