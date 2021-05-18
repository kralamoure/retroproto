// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedCommandClear struct{}

func (m BasicsAuthorizedCommandClear) ProtocolId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedCommandClear
}

func (m BasicsAuthorizedCommandClear) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedCommandClear) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
