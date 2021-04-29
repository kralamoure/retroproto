// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildLeaveTaxCollector struct{}

func (m GuildLeaveTaxCollector) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildLeaveTaxCollector
}

func (m GuildLeaveTaxCollector) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildLeaveTaxCollector) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
