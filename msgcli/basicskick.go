// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type BasicsKick struct{}

func (m BasicsKick) ProtocolId() d1proto.MsgCliId {
	return d1proto.BasicsKick
}

func (m BasicsKick) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *BasicsKick) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}