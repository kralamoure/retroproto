// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type TutorialCreate struct{}

func (m TutorialCreate) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.TutorialCreate
}

func (m TutorialCreate) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *TutorialCreate) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
