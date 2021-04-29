package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsNothing struct{}

func (m BasicsNothing) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.BasicsNothing
}

func (m BasicsNothing) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsNothing) Deserialize(extra string) error {
	return nil
}
