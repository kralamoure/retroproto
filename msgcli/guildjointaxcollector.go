// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildJoinTaxCollector struct{}

func (m GuildJoinTaxCollector) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildJoinTaxCollector
}

func (m GuildJoinTaxCollector) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildJoinTaxCollector) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}