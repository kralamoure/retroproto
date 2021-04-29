// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildGetInfosGeneral struct{}

func (m GuildGetInfosGeneral) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildGetInfosGeneral
}

func (m GuildGetInfosGeneral) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildGetInfosGeneral) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
