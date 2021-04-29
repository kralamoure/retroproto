// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestAreaAlignmentChanged struct{}

func (m ConquestAreaAlignmentChanged) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ConquestAreaAlignmentChanged
}

func (m ConquestAreaAlignmentChanged) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestAreaAlignmentChanged) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
