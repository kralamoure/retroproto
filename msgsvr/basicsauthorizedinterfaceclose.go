// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsAuthorizedInterfaceClose struct{}

func (m BasicsAuthorizedInterfaceClose) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.BasicsAuthorizedInterfaceClose
}

func (m BasicsAuthorizedInterfaceClose) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsAuthorizedInterfaceClose) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
