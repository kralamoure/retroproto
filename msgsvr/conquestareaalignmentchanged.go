// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ConquestAreaAlignmentChanged struct{}

func (m ConquestAreaAlignmentChanged) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestAreaAlignmentChanged
}

func (m ConquestAreaAlignmentChanged) MessageName() string {
	return "ConquestAreaAlignmentChanged"
}

func (m ConquestAreaAlignmentChanged) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestAreaAlignmentChanged) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
