// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildGetInfosBoosts struct{}

func (m GuildGetInfosBoosts) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildGetInfosBoosts
}

func (m GuildGetInfosBoosts) Serialized() (string, error) {
	return "", nil
}

func (m *GuildGetInfosBoosts) Deserialize(extra string) error {
	return nil
}
