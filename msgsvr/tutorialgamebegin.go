// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type TutorialGameBegin struct{}

func (m TutorialGameBegin) MessageId() retroproto.MsgSvrId {
	return retroproto.TutorialGameBegin
}

func (m TutorialGameBegin) MessageName() string {
	return "TutorialGameBegin"
}

func (m TutorialGameBegin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *TutorialGameBegin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
