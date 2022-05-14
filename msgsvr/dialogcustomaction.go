// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type DialogCustomAction struct{}

func (m DialogCustomAction) MessageId() retroproto.MsgSvrId {
	return retroproto.DialogCustomAction
}

func (m DialogCustomAction) MessageName() string {
	return "DialogCustomAction"
}

func (m DialogCustomAction) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *DialogCustomAction) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
