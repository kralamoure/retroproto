// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildJoinTaxCollector struct{}

func (m GuildJoinTaxCollector) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildJoinTaxCollector
}

func (m GuildJoinTaxCollector) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildJoinTaxCollector) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
