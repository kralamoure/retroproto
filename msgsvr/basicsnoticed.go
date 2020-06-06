package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type BasicsNoticed struct{}

func (m BasicsNoticed) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsNoticed
}

func (m BasicsNoticed) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsNoticed) Deserialize(extra string) error {
	return nil
}
