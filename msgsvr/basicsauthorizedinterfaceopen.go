// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsAuthorizedInterfaceOpen struct{}

func (m BasicsAuthorizedInterfaceOpen) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.BasicsAuthorizedInterfaceOpen
}

func (m BasicsAuthorizedInterfaceOpen) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsAuthorizedInterfaceOpen) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
