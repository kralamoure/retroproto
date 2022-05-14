// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type BasicsSanctionMe struct{}

func (m BasicsSanctionMe) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsSanctionMe
}

func (m BasicsSanctionMe) MessageName() string {
	return "BasicsSanctionMe"
}

func (m BasicsSanctionMe) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsSanctionMe) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
