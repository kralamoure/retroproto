// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type BasicsAuthorizedInterfaceOpen struct{}

func (m BasicsAuthorizedInterfaceOpen) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsAuthorizedInterfaceOpen
}

func (m BasicsAuthorizedInterfaceOpen) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsAuthorizedInterfaceOpen) Deserialize(extra string) error {
	return nil
}
