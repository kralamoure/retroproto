// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountServersListError struct{}

func (m AccountServersListError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountServersListError
}

func (m AccountServersListError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountServersListError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
