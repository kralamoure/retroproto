// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildRemoveTaxCollector struct{}

func (m GuildRemoveTaxCollector) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildRemoveTaxCollector
}

func (m GuildRemoveTaxCollector) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildRemoveTaxCollector) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
