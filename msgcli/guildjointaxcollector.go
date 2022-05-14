// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildJoinTaxCollector struct{}

func (m GuildJoinTaxCollector) MessageId() retroproto.MsgCliId {
	return retroproto.GuildJoinTaxCollector
}

func (m GuildJoinTaxCollector) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildJoinTaxCollector) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
