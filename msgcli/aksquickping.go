// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AksQuickPing struct{}

func (m AksQuickPing) ProtocolId() retroproto.MsgCliId {
	return retroproto.AksQuickPing
}

func (m AksQuickPing) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AksQuickPing) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
