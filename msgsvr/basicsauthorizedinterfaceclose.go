// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type BasicsAuthorizedInterfaceClose struct{}

func (m BasicsAuthorizedInterfaceClose) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsAuthorizedInterfaceClose
}

func (m BasicsAuthorizedInterfaceClose) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *BasicsAuthorizedInterfaceClose) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
