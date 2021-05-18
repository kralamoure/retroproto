// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type TutorialGameBegin struct{}

func (m TutorialGameBegin) ProtocolId() retroproto.MsgSvrId {
	return retroproto.TutorialGameBegin
}

func (m TutorialGameBegin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *TutorialGameBegin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
