// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildHireTaxCollector struct{}

func (m GuildHireTaxCollector) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildHireTaxCollector
}

func (m GuildHireTaxCollector) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildHireTaxCollector) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
