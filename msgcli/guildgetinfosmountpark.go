// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildGetInfosMountPark struct{}

func (m GuildGetInfosMountPark) ProtocolId() retroproto.MsgCliId {
	return retroproto.GuildGetInfosMountPark
}

func (m GuildGetInfosMountPark) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildGetInfosMountPark) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
