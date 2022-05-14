// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type BasicsWhoIs struct{}

func (m BasicsWhoIs) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsWhoIs
}

func (m BasicsWhoIs) MessageName() string {
	return "BasicsWhoIs"
}

func (m BasicsWhoIs) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsWhoIs) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
