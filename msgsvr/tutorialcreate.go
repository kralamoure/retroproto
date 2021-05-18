// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type TutorialCreate struct{}

func (m TutorialCreate) ProtocolId() retroproto.MsgSvrId {
	return retroproto.TutorialCreate
}

func (m TutorialCreate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *TutorialCreate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
