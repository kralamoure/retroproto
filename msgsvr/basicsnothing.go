package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsNothing struct{}

func (m BasicsNothing) ProtocolId() retroproto.MsgSvrId {
	return retroproto.BasicsNothing
}

func (m BasicsNothing) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsNothing) Deserialize(extra string) error {
	return nil
}
