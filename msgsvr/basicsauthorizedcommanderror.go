package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsAuthorizedCommandError struct{}

func (m BasicsAuthorizedCommandError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.BasicsAuthorizedCommandError
}

func (m BasicsAuthorizedCommandError) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsAuthorizedCommandError) Deserialize(extra string) error {
	return nil
}
