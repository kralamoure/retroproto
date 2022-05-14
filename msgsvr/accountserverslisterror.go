// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountServersListError struct{}

func (m AccountServersListError) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountServersListError
}

func (m AccountServersListError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountServersListError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
