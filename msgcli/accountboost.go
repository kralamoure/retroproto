// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountBoost struct{}

func (m AccountBoost) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountBoost
}

func (m AccountBoost) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *AccountBoost) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
