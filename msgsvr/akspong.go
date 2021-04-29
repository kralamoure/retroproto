package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AksPong struct{}

func (m AksPong) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AksPong
}

func (m AksPong) Serialized() (string, error) {
	return "", nil
}

func (m *AksPong) Deserialize(extra string) error {
	return nil
}
