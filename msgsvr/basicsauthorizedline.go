// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedLine struct{}

func (m BasicsAuthorizedLine) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedLine
}

func (m BasicsAuthorizedLine) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedLine) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
