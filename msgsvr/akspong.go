package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AksPong struct{}

func (m AksPong) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AksPong
}

func (m AksPong) Serialized() (string, error) {
	return "", nil
}

func (m *AksPong) Deserialize(extra string) error {
	return nil
}
