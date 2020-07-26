// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ConquestAreaAlignmentChanged struct{}

func (m ConquestAreaAlignmentChanged) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ConquestAreaAlignmentChanged
}

func (m ConquestAreaAlignmentChanged) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ConquestAreaAlignmentChanged) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
