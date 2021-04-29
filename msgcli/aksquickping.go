// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AksQuickPing struct{}

func (m AksQuickPing) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AksQuickPing
}

func (m AksQuickPing) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AksQuickPing) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
