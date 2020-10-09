// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type BasicsInvisible struct{}

func (m BasicsInvisible) ProtocolId() d1proto.MsgCliId {
	return d1proto.BasicsInvisible
}

func (m BasicsInvisible) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *BasicsInvisible) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
