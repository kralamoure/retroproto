package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type BasicsNothing struct{}

func (m BasicsNothing) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsNothing
}

func (m BasicsNothing) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsNothing) Deserialize(extra string) error {
	return nil
}
