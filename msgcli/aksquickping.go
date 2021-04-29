// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AksQuickPing struct{}

func (m AksQuickPing) ProtocolId() d1proto.MsgCliId {
	return d1proto.AksQuickPing
}

func (m AksQuickPing) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *AksQuickPing) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
