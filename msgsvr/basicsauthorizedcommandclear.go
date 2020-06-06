// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type BasicsAuthorizedCommandClear struct{}

func (m BasicsAuthorizedCommandClear) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsAuthorizedCommandClear
}

func (m BasicsAuthorizedCommandClear) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsAuthorizedCommandClear) Deserialize(extra string) error {
	return nil
}
