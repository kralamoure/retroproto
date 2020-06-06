// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type TutorialGameBegin struct{}

func (m TutorialGameBegin) ProtocolId() d1proto.MsgSvrId {
	return d1proto.TutorialGameBegin
}

func (m TutorialGameBegin) Serialized() (string, error) {
	return "", nil
}

func (m *TutorialGameBegin) Deserialize(extra string) error {
	return nil
}
