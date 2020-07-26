// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildRemoveTaxCollector struct{}

func (m GuildRemoveTaxCollector) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildRemoveTaxCollector
}

func (m GuildRemoveTaxCollector) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildRemoveTaxCollector) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
