// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountServersListError struct{}

func (m AccountServersListError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountServersListError
}

func (m AccountServersListError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *AccountServersListError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
