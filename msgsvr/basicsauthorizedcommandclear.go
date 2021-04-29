// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsAuthorizedCommandClear struct{}

func (m BasicsAuthorizedCommandClear) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.BasicsAuthorizedCommandClear
}

func (m BasicsAuthorizedCommandClear) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsAuthorizedCommandClear) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
