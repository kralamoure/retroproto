// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsAway struct{}

func (m BasicsAway) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.BasicsAway
}

func (m BasicsAway) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsAway) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
