// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildInfosMountPark struct{}

func (m GuildInfosMountPark) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildInfosMountPark
}

func (m GuildInfosMountPark) Serialized() (string, error) {
	return "", nil
}

func (m *GuildInfosMountPark) Deserialize(extra string) error {
	return nil
}
