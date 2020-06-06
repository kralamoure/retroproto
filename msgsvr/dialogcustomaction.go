// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type DialogCustomAction struct{}

func (m DialogCustomAction) ProtocolId() d1proto.MsgSvrId {
	return d1proto.DialogCustomAction
}

func (m DialogCustomAction) Serialized() (string, error) {
	return "", nil
}

func (m *DialogCustomAction) Deserialize(extra string) error {
	return nil
}
