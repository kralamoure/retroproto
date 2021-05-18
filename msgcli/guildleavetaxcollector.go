// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildLeaveTaxCollector struct{}

func (m GuildLeaveTaxCollector) ProtocolId() retroproto.MsgCliId {
	return retroproto.GuildLeaveTaxCollector
}

func (m GuildLeaveTaxCollector) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildLeaveTaxCollector) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
