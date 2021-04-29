// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type BasicsAuthorizedLine struct{}

func (m BasicsAuthorizedLine) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsAuthorizedLine
}

func (m BasicsAuthorizedLine) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *BasicsAuthorizedLine) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
