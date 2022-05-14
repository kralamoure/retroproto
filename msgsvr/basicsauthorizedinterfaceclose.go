// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedInterfaceClose struct{}

func (m BasicsAuthorizedInterfaceClose) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedInterfaceClose
}

func (m BasicsAuthorizedInterfaceClose) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedInterfaceClose) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
