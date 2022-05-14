// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountRescue struct{}

func (m AccountRescue) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountRescue
}

func (m AccountRescue) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountRescue) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
