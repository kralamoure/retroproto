// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountNewLevel struct{}

func (m AccountNewLevel) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountNewLevel
}

func (m AccountNewLevel) Serialized() (string, error) {
	return "", nil
}

func (m *AccountNewLevel) Deserialize(extra string) error {
	return nil
}
