// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type BasicsWhoIs struct{}

func (m BasicsWhoIs) ProtocolId() d1proto.MsgCliId {
	return d1proto.BasicsWhoIs
}

func (m BasicsWhoIs) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *BasicsWhoIs) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}