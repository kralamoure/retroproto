// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsAuthorizedCommandSuccess struct{}

func (m BasicsAuthorizedCommandSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.BasicsAuthorizedCommandSuccess
}

func (m BasicsAuthorizedCommandSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsAuthorizedCommandSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
