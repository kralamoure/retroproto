// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsSanctionMe struct{}

func (m BasicsSanctionMe) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.BasicsSanctionMe
}

func (m BasicsSanctionMe) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsSanctionMe) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
