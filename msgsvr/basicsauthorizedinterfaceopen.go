// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedInterfaceOpen struct{}

func (m BasicsAuthorizedInterfaceOpen) ProtocolId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedInterfaceOpen
}

func (m BasicsAuthorizedInterfaceOpen) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedInterfaceOpen) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
