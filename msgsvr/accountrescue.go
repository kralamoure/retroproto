// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountRescue struct{}

func (m AccountRescue) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountRescue
}

func (m AccountRescue) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountRescue) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
