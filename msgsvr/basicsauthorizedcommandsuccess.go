// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type BasicsAuthorizedCommandSuccess struct{}

func (m BasicsAuthorizedCommandSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsAuthorizedCommandSuccess
}

func (m BasicsAuthorizedCommandSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsAuthorizedCommandSuccess) Deserialize(extra string) error {
	return nil
}
