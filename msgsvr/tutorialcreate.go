// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type TutorialCreate struct{}

func (m TutorialCreate) ProtocolId() d1proto.MsgSvrId {
	return d1proto.TutorialCreate
}

func (m TutorialCreate) Serialized() (string, error) {
	return "", nil
}

func (m *TutorialCreate) Deserialize(extra string) error {
	return nil
}
