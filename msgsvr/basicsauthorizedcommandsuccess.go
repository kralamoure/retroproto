// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedCommandSuccess struct{}

func (m BasicsAuthorizedCommandSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedCommandSuccess
}

func (m BasicsAuthorizedCommandSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedCommandSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
