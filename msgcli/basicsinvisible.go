// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsInvisible struct{}

func (m BasicsInvisible) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.BasicsInvisible
}

func (m BasicsInvisible) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsInvisible) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
