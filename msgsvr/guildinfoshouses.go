// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildInfosHouses struct{}

func (m GuildInfosHouses) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildInfosHouses
}

func (m GuildInfosHouses) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildInfosHouses) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
