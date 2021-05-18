// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type BasicsSanctionMe struct{}

func (m BasicsSanctionMe) ProtocolId() retroproto.MsgCliId {
	return retroproto.BasicsSanctionMe
}

func (m BasicsSanctionMe) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsSanctionMe) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
