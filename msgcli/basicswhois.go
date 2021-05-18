// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type BasicsWhoIs struct{}

func (m BasicsWhoIs) ProtocolId() retroproto.MsgCliId {
	return retroproto.BasicsWhoIs
}

func (m BasicsWhoIs) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsWhoIs) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
