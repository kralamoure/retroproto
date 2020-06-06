// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type DialogQuestion struct{}

func (m DialogQuestion) ProtocolId() d1proto.MsgSvrId {
	return d1proto.DialogQuestion
}

func (m DialogQuestion) Serialized() (string, error) {
	return "", nil
}

func (m *DialogQuestion) Deserialize(extra string) error {
	return nil
}
