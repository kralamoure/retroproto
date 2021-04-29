// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildGetInfosGeneral struct{}

func (m GuildGetInfosGeneral) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildGetInfosGeneral
}

func (m GuildGetInfosGeneral) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildGetInfosGeneral) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
