// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedCommandClear struct{}

func (m BasicsAuthorizedCommandClear) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedCommandClear
}

func (m BasicsAuthorizedCommandClear) MessageName() string {
	return "BasicsAuthorizedCommandClear"
}

func (m BasicsAuthorizedCommandClear) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedCommandClear) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
