// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildGetInfosBoosts struct{}

func (m GuildGetInfosBoosts) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildGetInfosBoosts
}

func (m GuildGetInfosBoosts) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildGetInfosBoosts) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
