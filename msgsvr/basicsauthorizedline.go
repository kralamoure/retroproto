// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsAuthorizedLine struct{}

func (m BasicsAuthorizedLine) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.BasicsAuthorizedLine
}

func (m BasicsAuthorizedLine) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsAuthorizedLine) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
