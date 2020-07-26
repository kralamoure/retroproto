// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildHireTaxCollector struct{}

func (m GuildHireTaxCollector) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildHireTaxCollector
}

func (m GuildHireTaxCollector) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildHireTaxCollector) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
