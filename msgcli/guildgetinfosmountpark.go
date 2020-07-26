// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildGetInfosMountPark struct{}

func (m GuildGetInfosMountPark) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildGetInfosMountPark
}

func (m GuildGetInfosMountPark) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildGetInfosMountPark) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
