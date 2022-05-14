package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AksPong struct{}

func (m AksPong) MessageId() retroproto.MsgSvrId {
	return retroproto.AksPong
}

func (m AksPong) Serialized() (string, error) {
	return "", nil
}

func (m *AksPong) Deserialize(extra string) error {
	return nil
}
