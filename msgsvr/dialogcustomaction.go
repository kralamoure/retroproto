// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type DialogCustomAction struct{}

func (m DialogCustomAction) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.DialogCustomAction
}

func (m DialogCustomAction) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *DialogCustomAction) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
