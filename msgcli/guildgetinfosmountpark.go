// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildGetInfosMountPark struct{}

func (m GuildGetInfosMountPark) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildGetInfosMountPark
}

func (m GuildGetInfosMountPark) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildGetInfosMountPark) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
