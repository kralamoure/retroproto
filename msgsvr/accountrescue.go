// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountRescue struct{}

func (m AccountRescue) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountRescue
}

func (m AccountRescue) Serialized() (string, error) {
	return "", nil
}

func (m *AccountRescue) Deserialize(extra string) error {
	return nil
}
