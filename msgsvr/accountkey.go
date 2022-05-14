// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountKey struct{}

func (m AccountKey) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountKey
}

func (m AccountKey) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountKey) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
