// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AksQuickPing struct{}

func (m AksQuickPing) MessageId() retroproto.MsgCliId {
	return retroproto.AksQuickPing
}

func (m AksQuickPing) MessageName() string {
	return "AksQuickPing"
}

func (m AksQuickPing) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AksQuickPing) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
