package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type BasicsAuthorizedCommandError struct{}

func (m BasicsAuthorizedCommandError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsAuthorizedCommandError
}

func (m BasicsAuthorizedCommandError) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsAuthorizedCommandError) Deserialize(extra string) error {
	return nil
}
