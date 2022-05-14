// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildRemoveTaxCollector struct{}

func (m GuildRemoveTaxCollector) MessageId() retroproto.MsgCliId {
	return retroproto.GuildRemoveTaxCollector
}

func (m GuildRemoveTaxCollector) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildRemoveTaxCollector) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
