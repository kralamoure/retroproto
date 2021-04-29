// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsWhoIs struct{}

func (m BasicsWhoIs) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.BasicsWhoIs
}

func (m BasicsWhoIs) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsWhoIs) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
